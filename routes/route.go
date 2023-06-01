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
	e.PUT("/user/:id", controller.UpdateUserController)
	e.DELETE("/user/:id", controller.DeleteUserController)
	e.GET("/user", controller.GetUsersController)
	//eJwt := e.Group("/jwt")

	NewThreadControllers(e)
	e.Logger.Fatal(e.Start(":8000"))
	return e
}

func NewThreadControllers(e *echo.Echo) {
	e.GET("threads", controller.GetThreadController)
	e.GET("threads/:id", controller.GetThreadsIDController)
	e.POST("threads", controller.CreateThreadsController)
	e.DELETE("threads/:id", controller.DeleteThreadsController)
	e.PUT("threads/:id", controller.UpdateThreadsController)
}
