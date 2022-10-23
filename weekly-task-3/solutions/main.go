package main

import (
	"github.com/arvinpaundra/golang-blog/config"
	"github.com/arvinpaundra/golang-blog/database"
	"github.com/arvinpaundra/golang-blog/route"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitConfig()
	database.InitDatabase()

	e := echo.New()

	route.New(e, database.DB)

	e.Logger.Fatal(e.Start(config.Cfg.APIPort))
}
