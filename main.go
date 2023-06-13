package main

import (
	"Capstone/database"
	"Capstone/routes"

)


func main() {
	database.InitDB()
	e := routes.New()
	e.Start("8000")

}
