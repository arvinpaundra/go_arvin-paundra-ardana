package main

import (
	"os"

	"github.com/arvinpaundra/unit-test/config"
	"github.com/arvinpaundra/unit-test/middlewares"
	"github.com/arvinpaundra/unit-test/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	middlewares.LogMiddlewares(e)

	PORT := os.Getenv("APIPort")

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(PORT))
}
