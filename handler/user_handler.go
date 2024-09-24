package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Ryan",
		"email": "edwardlauv@gmail.com",
	})
}
func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string `json:"email"`
		FullName string `json:"full_name"`
	}
	user := User{
		Email:    "edwardlauv@gmail.com",
		FullName: "lauv",
	}
	return c.JSON(http.StatusOK, user)
}
