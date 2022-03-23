package user_route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-vote/handler/user_handler"
	"go-vote/infra"
	"go-vote/util/jwt"
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
	user.GET("/profile/:id", user_handler.GetProfile, middleware.JWTWithConfig(jwt.GetAccessConfig()))

	// ~/user/auth group
	auth := user.Group(groupAuth)
	auth.POST("/login", user_handler.Login)
	auth.POST("/refresh", user_handler.Refresh, middleware.JWTWithConfig(jwt.GetRefreshConfig()))
}
