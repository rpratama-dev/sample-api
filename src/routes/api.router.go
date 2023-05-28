package routes

import (
	"github.com/labstack/echo/v4"
	v1Routes "github.com/rpratama-dev/sample-api/src/routes/v1"
	v2Routes "github.com/rpratama-dev/sample-api/src/routes/v2"
)

func ApiRouter(router *echo.Group) {
	v1Routes.RegisterRoutes(router.Group("/v1"))
	v2Routes.RegisterRoutes(router.Group("/v2"))
}
