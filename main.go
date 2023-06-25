package main

import (
	"Capstone/database"
	"Capstone/routes"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "8080"

func main() {
	database.Init()

	app := echo.New()

	routes.New(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	appPort := fmt.Sprintf(":%s", port)

	app.Logger.Fatal(app.Start(appPort))

}
