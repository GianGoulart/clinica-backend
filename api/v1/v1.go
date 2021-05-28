package v1

import (
	"github.com/GianGoulart/Clinica_backend/api/middleware"
	"github.com/GianGoulart/Clinica_backend/api/v1/health"
	"github.com/GianGoulart/Clinica_backend/api/v1/item"
	"github.com/GianGoulart/Clinica_backend/app"
	"github.com/labstack/echo/v4"
)

// Register regristra as rotas v1
func Register(g *echo.Group, apps *app.Container, middleware *middleware.Middleware) {
	v1 := g.Group("/v1", middleware.Session.InjectSession)

	health.Register(v1.Group("/health"), apps, middleware)
	item.Register(v1.Group("/item"), apps, middleware)
}
