package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-vote/infra"
	"go-vote/route"
)

func main() {
	// set log config
	log.SetPrefix("go-vote")

	// load infra
	inf, err := infra.Init()
	if err != nil {
		log.Errorf("failed to load infrastructure: %v", err)
		return
	}

	// init echo service
	e := echo.New()
	route.Init(e, inf)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", inf.Config.ServicePort)))
}
