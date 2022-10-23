package main

import (
	"github.com/arvinpaundra/clean-architecture/config"
	"github.com/arvinpaundra/clean-architecture/database"
	"github.com/arvinpaundra/clean-architecture/route"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitConfig()
	database.InitDatabase()

	e := echo.New()

	route.New(e, database.DB)

	e.Logger.Fatal(e.Start(config.Cfg.APIPort))
}
