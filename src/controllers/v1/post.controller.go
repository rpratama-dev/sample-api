package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexPost(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v1 posts index")
}

func StorePost(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v1 post store")
}

func ShowPost(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v1 post show")
}

func DestroyPost(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v1 post destroy")
}
