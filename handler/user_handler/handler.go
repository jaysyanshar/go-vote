package user_handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"go-vote/model"
	"go-vote/repository/user_repository"
	"go-vote/service/user_service"
)

var service user_service.UserService

func Init(db *sql.DB) {
	repo := user_repository.Init(db)
	service = user_service.Init(&repo)
}

func Register(c echo.Context) error {
	req := &model.RegisterUserReq{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Error: err.Error()})
	}
	res, err := service.Register(req)
	if err != nil {
		return c.JSON(res.Status, model.ResponseError{Error: err.Error()})
	}
	return c.JSON(res.Status, res)
}
