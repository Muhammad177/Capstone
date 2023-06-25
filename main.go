package main

import (
	"Capstone/database"
	"Capstone/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "8080"

func main() {
	if env := os.Getenv("ENV"); env != "production" {
		log.Println(env)
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("failed to load .env")
		}
	}
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
