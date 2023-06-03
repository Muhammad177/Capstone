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
	e.POST("/login", controller.LoginController)
	e.POST("/user", controller.CreateUserController)
	eJwt := e.Group("/jwt")
	eJwt.PUT("/admin/:id", controller.UpdateUserAdminController)
	eJwt.DELETE("/admin/:id", controller.DeleteUserAdminController)
	eJwt.GET("/admin", controller.GetUsersAdminController)
	eJwt.GET("/admin/:id", controller.GetUserByidAdminController)
	eJwt.PUT("/user/:id", controller.UpdateUserController)
	eJwt.DELETE("/user/:id", controller.DeleteUserController)
	eJwt.GET("/user", controller.GetUserController)
	//confirm
	e.Logger.Fatal(e.Start(":8000"))
	return e
}
