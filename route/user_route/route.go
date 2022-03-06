package user_route

import (
	"go-vote/infra"

	"github.com/labstack/echo/v4"

	"go-vote/handler/user_handler"
)

const (
	groupPrefix = "users"
)

func Init(e *echo.Echo, inf *infra.Infra) {
	user_handler.Init(inf)
	group := e.Group(groupPrefix)
	group.POST("/register", user_handler.Register)
	group.GET("/profile/:id", user_handler.GetProfile)
}
