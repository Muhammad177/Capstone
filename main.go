package main

import (
	"Capstone/database"
	"Capstone/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env")
	}
	database.InitDB()
	e := routes.New()
	e.Start("8000")

}
