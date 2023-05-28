package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group) {
	UserRouter(g.Group("/users"));
	PostRouter(g.Group("/posts"));
}
