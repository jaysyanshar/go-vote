package route

import (
	"database/sql"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"go-vote/route/user_route"
)

func Init(e *echo.Echo, db *sql.DB) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	user_route.Init(e, db)
}
