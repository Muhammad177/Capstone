package main

import (
	"Capstone/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if env := os.Getenv("ENV"); env != "production" {
		log.Println(env)
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("failed to load .env")
		}
	}
	database.Init()

}
