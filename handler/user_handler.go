package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"trending-github-golang/log"
	"trending-github-golang/model"
	req "trending-github-golang/model/req"
	"trending-github-golang/responsitory/repo_impl"
	"trending-github-golang/security"
)

type UserHandler struct {
	UserRepoImpl repo_impl.UserRepoImpl
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Ryan",
		"email": "edwardlauv@gmail.com",
	})
}
func (u *UserHandler) HandleSignUp(c echo.Context) error {
	// check request binding
	request := req.ReqSignUp{}
	if err := c.Bind(&request); err != nil {
		log.Error(err.Error())
		return HandleError(err, c)
	}
	// validate model request
	//validate := validator.New()
	if err := c.Validate(request); err != nil {
		log.Error(err.Error())
		return HandleError(err, c)
	}
	// hash password
	hash, err := security.HashPassword(request.Password)
	if err != nil {
		log.Error(err.Error())
		return HandleError(err, c)
	}
	role := model.Role(model.RoleMember).String()
	userid, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return HandleError(err, c)
	}
	user := model.User{
		Email:    request.Email,
		FullName: request.FullName,
		Role:     role,
		Password: hash,
		Userid:   userid.String(),
	}
	user, err = u.UserRepoImpl.SaveUser(c.Request().Context(), user)
	if err != nil {
		log.Error(err.Error())
		return HandleError(err, c)
	}
	// import to db -> use user repo to run this
	return c.JSON(http.StatusOK, user)
}

func HandleError(err error, c echo.Context) error {
	return c.JSON(http.StatusBadRequest, model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
		Data:       nil,
	})
}
