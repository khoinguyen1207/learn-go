package app

import (
	"log"
	"project-shopping/internal/config"
	"project-shopping/internal/routes"
	"project-shopping/internal/validation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()

	if err := validation.InitValidator(); err != nil {
		log.Fatal("Failed to initialize validator:", err)
	}

	loadEnv()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config:  cfg,
		router:  r,
		modules: modules,
	}
}

func (app *Application) Run() error {
	return app.router.Run(app.config.Port)
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))

	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}
