package user_route

import (
	"go-vote/infra"

	"github.com/labstack/echo/v4"

	"go-vote/handler/user_handler"
)

const (
	groupUser = "/user"
	groupAuth = "/auth"
)

func Init(e *echo.Echo, inf *infra.Infra) {
	user_handler.Init(inf)

	// ~/user group
	user := e.Group(groupUser)
	user.POST("/register", user_handler.Register)
	user.GET("/profile/:id", user_handler.GetProfile)

	// ~/user/auth group
	auth := user.Group(groupAuth)
	auth.POST("/login", user_handler.Login)
}
