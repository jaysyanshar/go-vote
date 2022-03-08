package user_handler

import (
	"github.com/labstack/echo/v4"
	"go-vote/infra"
	"go-vote/model"
	"go-vote/repository/user_repository"
	"go-vote/service/user_service"
	"go-vote/util/convert"
	"go-vote/util/response"
	"net/http"
)

var service user_service.UserService

func Init(inf *infra.Infra) {
	repo := user_repository.Init(inf)
	service = user_service.Init(&repo)
}

func Register(c echo.Context) error {
	req := &model.RegisterUserReq{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.MakeErrorResponse(err.Error()))
	}
	res, err := service.Register(req)
	if err != nil {
		return c.JSON(res.Status, response.MakeErrorResponse(err.Error()))
	}
	return c.JSON(res.Status, res)
}

func GetProfile(c echo.Context) error {
	param := c.Param("id")
	id, err := convert.StrToInt64(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MakeErrorResponse(err.Error()))
	}
	res, err := service.GetProfile(id)
	if err != nil {
		return c.JSON(res.Status, response.MakeErrorResponse(err.Error()))
	}
	return c.JSON(res.Status, res)
}

func Login(c echo.Context) error {
	username := c.Request().Header.Get("Username")
	password := c.Request().Header.Get("Password")
	req := &model.LoginUserReq{
		Email:    username,
		Password: password,
	}
	res, err := service.Login(req)
	if err != nil {
		return c.JSON(res.Status, response.MakeErrorResponse(err.Error()))
	}
	return c.JSON(res.Status, res)
}
