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
	eJwt := e.Group("/jwt")
	eJwt.PUT("/user/:id", controller.UpdateUserController)
	eJwt.DELETE("/user/:id", controller.DeleteUserController)
	eJwt.GET("/user", controller.GetUsersController)
	eJwt.GET("/user:id", controller.GetUserByidController)
	e.Logger.Fatal(e.Start(":8000"))
	return e
}
