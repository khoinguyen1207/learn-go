package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project-shopping/internal/config"
	"project-shopping/internal/db"
	"project-shopping/internal/db/sqlc"
	"project-shopping/internal/routes"
	"project-shopping/internal/validation"
	"project-shopping/pkg/auth"
	"project-shopping/pkg/cache"
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
		log.Fatal("Failed to initialize validator:", err)
	}

	if err := db.InitDB(cfg); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	redisClient := config.InitRedis(cfg)
	cacheService := cache.NewCacheService(redisClient)
	jwtService := auth.NewJWTService(&cfg.Jwt)

	moduleContext := &ModuleContext{
		db:    db.GetDB(),
		cache: cacheService,
		jwt:   jwtService,
	}

	modules := []Module{
		NewUserModule(moduleContext),
		NewAuthModule(moduleContext),
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
		log.Printf("✅ Server is running on port %s", app.config.Port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	<-quit
	log.Println("⚠️  Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	log.Println("✅ Stopped server gracefully")

	return nil
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))

	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}
