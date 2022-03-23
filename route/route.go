package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-vote/infra"
	"go-vote/route/user_route"
)

func Init(e *echo.Echo, inf *infra.Infra) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	user_route.Init(e, inf)
}
