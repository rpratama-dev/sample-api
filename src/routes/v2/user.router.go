package routes

import (
	echo "github.com/labstack/echo/v4"
	"github.com/rpratama-dev/sample-api/src/controllers/v2"
)

func UserRouter(router *echo.Group) {
	router.GET("", controllers.IndexUser)
	router.GET("/:id", controllers.ShowUser)
	router.POST("", controllers.StoreUser)
	router.PUT("/:id", controllers.UpdateUser)
}
