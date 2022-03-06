package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-vote/route"

	"go-vote/config"
	sql "go-vote/infra/db"
)

func main() {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Errorf("failed to load config: %v", err)
		return
	}

	// load db
	db, err := sql.LoadDb(*cfg)
	if err != nil {
		log.Errorf("failed to load db: %v", err)
		return
	}

	// init echo service
	e := echo.New()
	route.Init(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServicePort)))
}
