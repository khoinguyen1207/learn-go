package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"project-shopping/internal/config"
	"project-shopping/internal/db"
	"project-shopping/internal/db/sqlc"
	"project-shopping/internal/routes"
	"project-shopping/internal/utils"
	"project-shopping/internal/validation"
	"project-shopping/pkg/auth"
	"project-shopping/pkg/cache"
	"project-shopping/pkg/logger"
	"project-shopping/pkg/mail"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

type ModuleContext struct {
	db    sqlc.Querier
	cache cache.CacheService
	jwt   auth.JWTService
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()

	if err := validation.InitValidator(); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to initialize validator")
	}

	if err := db.InitDB(cfg); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	redisClient := config.InitRedis(cfg)
	cacheService := cache.NewCacheService(redisClient)
	jwtService := auth.NewJWTService(cfg)
	mailProvider, err := mail.NewMailProvider(cfg)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to initialize mail provider")
	}
	mailLogger := utils.NewLoggerWithPath("mail.log", "info")
	mailService := mail.NewMailService(cfg, mailProvider, mailLogger)

	moduleContext := &ModuleContext{
		db:    db.GetDB(),
		cache: cacheService,
		jwt:   jwtService,
	}

	modules := []Module{
		NewUserModule(moduleContext),
		NewAuthModule(moduleContext, mailService),
	}

	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config:  cfg,
		router:  r,
		modules: modules,
	}
}

func (app *Application) Run() error {
	srv := &http.Server{
		Addr:    app.config.Port,
		Handler: app.router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		logger.Log.Info().Msgf("Server is running on port %s", app.config.Port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	<-quit
	logger.Log.Info().Msg("⚠️ Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Error().Err(err).Msg("❌ Server forced to shutdown")
	}

	logger.Log.Info().Msg("✅ Stopped server gracefully")

	return nil
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))

	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}
