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
	e.GET("/threads", controller.GetThreadControllerByTitle)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.PUT("/admin/:id", controller.UpdateUserAdminController)
	eJwt.DELETE("/admin/:id", controller.DeleteUserAdminController)
	eJwt.GET("/admin", controller.GetUsersAdminController)
	eJwt.GET("/admin/:id", controller.GetUserByidAdminController)
	eJwt.PUT("/user", controller.UpdateUserController)
	eJwt.DELETE("/user", controller.DeleteUserController)
	eJwt.GET("/user", controller.GetUserController)
	eJwt.GET("/image", controller.GetImageHandler)
	//confirm
	NewThreadControllers(eJwt)

	e.Logger.Fatal(e.Start(":8000"))
	return e
}

func NewThreadControllers(e *echo.Group) {
	e.GET("/admin/threads", controller.GetThreadController)
	e.GET("/threads/:id", controller.GetThreadsIDController)
	e.POST("/threads", controller.CreateThreadsController)
	e.DELETE("/admin/threads/:id", controller.DeleteThreadsControllerAdmin)
	e.DELETE("/threads/:id", controller.DeleteThreadsControllerAdmin)
	e.PUT("/admin/threads/:id", controller.UpdateThreadsControllerAdmin)
	e.PUT("/threads/:id", controller.UpdateThreadsControllerAdmin)
}
