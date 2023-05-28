package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	appMiddleware "github.com/rpratama-dev/sample-api/src/middleware"
	"github.com/rpratama-dev/sample-api/src/routes"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	// Register router API
	routes.ApiRouter(e.Group("/api"))
	// Custom middleware to handle router not found error
	e.Use(appMiddleware.RouterErrorMiddleware);
	e.Logger.Fatal(e.Start(":6090"))
}
