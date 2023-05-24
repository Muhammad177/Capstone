package routes

import (
	"Capstone/controller"
	"Capstone/midleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	midleware.LogMiddleware(e)
	// routing with query parameter
	e.POST("/login", controller.LoginUserController)
	e.POST("/user", controller.CreateUserController)
	//eJwt := e.Group("/jwt")
	e.Logger.Fatal(e.Start(":8000"))
	return e
}
