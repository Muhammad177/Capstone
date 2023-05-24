package main

import (
	"Remidi/database"
	"Remidi/routes"

)


func main() {
	database.InitDB()
	e := routes.New()
	e.Start("8000")

}
