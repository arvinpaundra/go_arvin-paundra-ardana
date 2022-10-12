package main

import (
	"os"

	"github.com/arvinpaundra/weekly-task-2/configs"
	"github.com/arvinpaundra/weekly-task-2/middlewares"
	"github.com/arvinpaundra/weekly-task-2/routes"
)

func main() {
	configs.InitDB()

	e := routes.New()

	middlewares.LoggingMiddlewares(e)

	PORT := os.Getenv("APIPort")

	e.Logger.Fatal(e.Start(PORT))
}
