package main

import (
	"os"

	"github.com/arvinpaundra/rest-orm/config"
	"github.com/arvinpaundra/rest-orm/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	PORT := os.Getenv("APIPort")

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(PORT))
}
