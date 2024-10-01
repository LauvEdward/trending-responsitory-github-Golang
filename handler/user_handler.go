package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"trending-github-golang/log"
	"trending-github-golang/model"
	req "trending-github-golang/model/req"
	"trending-github-golang/responsitory"
	"trending-github-golang/security"
)

type UserHandler struct {
	UserRepoImpl responsitory.UserRepo
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	request := req.ReqSignIn{}
	if err := c.Bind(&request); err != nil {
		log.Error(err.Error())
		return HandleError(err, c)
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		log.Error(err.Error())
		return HandleError(err, c)
	}
	user, err := u.UserRepoImpl.CheckUser(c.Request().Context(), request)
	if err != nil {
		log.Error(err.Error())
		return HandleError(err, c)
	}
	token, err := security.GenerateToken(user)
	if err != nil {
		log.Info(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data:       user,
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
	validate := validator.New()
	//log.Info(request)
	if err := validate.Struct(request); err != nil {
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
	useridString := userid.String()
	user := model.User{
		Email:    request.Email,
		FullName: request.FullName,
		Role:     role,
		Password: hash,
		UserId:   useridString,
	}
	user, err = u.UserRepoImpl.SaveUser(c.Request().Context(), user)
	if err != nil {
		log.Info(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	token, err := security.GenerateToken(user)
	if err != nil {
		log.Info(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token
	// import to db -> use user repo to run this
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "User created successfully",
		Data:       user,
	})
}

func HandleError(err error, c echo.Context) error {
	return c.JSON(http.StatusBadRequest, model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
		Data:       nil,
	})
}
