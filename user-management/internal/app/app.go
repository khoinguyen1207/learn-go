package app

import (
	"user-management/internal/config"
	"user-management/internal/routes"

	"github.com/gin-gonic/gin"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getmModuleRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
	}
}

func (app *Application) Run() error {
	return app.router.Run(app.config.Port)
}

func getmModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))

	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}
