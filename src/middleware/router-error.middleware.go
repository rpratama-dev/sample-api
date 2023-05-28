package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouterErrorMiddleware (next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		// Check if the error is a router not found error
		if err != nil && err == echo.ErrNotFound {
			// Handle the router not found error
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "Route not found",
			})
		}

		return err
	}
}
