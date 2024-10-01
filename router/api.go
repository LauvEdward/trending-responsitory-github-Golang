package router

import (
	"github.com/labstack/echo/v4"
	"trending-github-golang/handler"
)

type API struct {
	Echo        *echo.Echo
	UserHandler *handler.UserHandler
}

func (a *API) SetUp() {
	g := a.Echo.Group("/user")
	g.GET("/sign-in", a.UserHandler.HandleSignIn)
	g.GET("/sign-up", a.UserHandler.HandleSignUp)
}
