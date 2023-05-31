package main

import (
	"thread/config"
	"thread/database"
	"thread/route"
)

func main() {
	config.InitConfig()

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	if err := database.MigrateDB(db); err != nil {
		panic(err)
	}

	e := route.InitRoute(db)

	e.Logger.Fatal(e.Start(config.Cfg.API_PORT))
}
