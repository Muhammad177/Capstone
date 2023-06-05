package routes

import (
	"Capstone/constant"
	"Capstone/controller"
	"Capstone/midleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	midleware.LogMiddleware(e)
	// routing with query parameter
	e.POST("/login", controller.LoginController)
	e.POST("/login/admin", controller.LoginAdminController)
	e.POST("/user", controller.CreateUserController)
	e.GET("/image/:uuid", controller.GetImageHandler)
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.PUT("/admin/:id", controller.UpdateUserAdminController)
	eJwt.DELETE("/admin/:id", controller.DeleteUserAdminController)
	eJwt.GET("/admin", controller.GetUsersAdminController)
	eJwt.GET("/admin/:id", controller.GetUserByidAdminController)
	eJwt.PUT("/user/:id", controller.UpdateUserController)
	eJwt.DELETE("/user/:id", controller.DeleteUserController)
	eJwt.GET("/user", controller.GetUserController)
	//confirm
	NewThreadControllers(eJwt)

	e.Logger.Fatal(e.Start(":8000"))
	return e
}

func NewThreadControllers(e *echo.Group) {
	e.GET("/threads", controller.GetThreadController)
	e.GET("/threads/:id", controller.GetThreadsIDController)
	e.POST("/threads", controller.CreateThreadsController)
	e.DELETE("/threads/:id", controller.DeleteThreadsController)
	e.PUT("/threads/:id", controller.UpdateThreadsController)
}
