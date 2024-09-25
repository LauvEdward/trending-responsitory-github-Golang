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
	a.Echo.GET("/user/sign-in", a.UserHandler.HandleSignIn)
	a.Echo.GET("/user/sign-up", a.UserHandler.HandleSignUp)
}
