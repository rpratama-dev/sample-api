package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v1 users index")
}

func StoreUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v1 user store")
}

func ShowUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v1 user show")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v1 user update")
}
