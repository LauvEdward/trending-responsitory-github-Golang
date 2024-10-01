package Middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MiddleWares() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, MiddleWares!")
		}
	}
}
