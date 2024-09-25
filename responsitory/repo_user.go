package responsitory

import (
	"github.com/labstack/echo/v4"
	"trending-github-golang/model"
)

type UserRepo interface{}

func SaveUser(context echo.Context, user model.User) {}
