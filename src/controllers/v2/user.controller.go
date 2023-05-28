package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v2 users index")
}

func StoreUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v2 user store")
}

func ShowUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v2 user show")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is the v2 user update")
}
