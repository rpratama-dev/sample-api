package routes

import (
	echo "github.com/labstack/echo/v4"
	controllers "github.com/rpratama-dev/sample-api/src/controllers/v1"
)

func PostRouter(router *echo.Group) {
	router.GET("", controllers.IndexPost)
	router.GET("/:id", controllers.ShowPost)
	router.POST("", controllers.StorePost)
	router.DELETE("/:id", controllers.DestroyPost)
}
