package user_handler

import (
	"github.com/labstack/echo/v4"
	"go-vote/infra"
	"go-vote/model"
	"go-vote/repository/auth_repository"
	"go-vote/repository/user_repository"
	"go-vote/service/user_service"
	"go-vote/util/constant"
	"go-vote/util/convert"
	"go-vote/util/response"
	"net/http"
)

var service user_service.UserService

func Init(inf *infra.Infra) {
	userRepo := user_repository.Init(inf)
	authRepo := auth_repository.Init(inf)
	service = user_service.Init(&userRepo, &authRepo)
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
	param := c.Param(constant.HttpPathId)
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
	username := c.Request().Header.Get(constant.HttpHeaderUsername)
	password := c.Request().Header.Get(constant.HttpHeaderPassword)
	ip := c.RealIP()
	req := &model.LoginUserReq{
		Email:     username,
		Password:  password,
		IpAddress: ip,
	}
	res, err := service.Login(req)
	if err != nil {
		return c.JSON(res.Status, response.MakeErrorResponse(err.Error()))
	}
	refresh := &http.Cookie{Name: constant.HttpCookieRefreshToken, Value: res.RefreshToken}
	c.SetCookie(refresh)
	return c.JSON(res.Status, res)
}

func Refresh(c echo.Context) error {
	refresh, _ := c.Cookie(constant.HttpCookieRefreshToken)
	ip := c.RealIP()
	req := &model.RefreshUserReq{
		RefreshToken: refresh.Value,
		IpAddress:    ip,
	}
	res, err := service.Refresh(req)
	if err != nil {
		return c.JSON(res.Status, response.MakeErrorResponse(err.Error()))
	}
	refresh.Value = res.RefreshToken
	c.SetCookie(refresh)
	return c.JSON(res.Status, res)
}
