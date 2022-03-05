package user_route

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"go-vote/handler/user_handler"
)

const (
	groupPrefix = "users"
)

func Init(e *echo.Echo, db *sql.DB) {
	user_handler.Init(db)
	group := e.Group(groupPrefix)
	group.POST("/register", user_handler.Register)
}
