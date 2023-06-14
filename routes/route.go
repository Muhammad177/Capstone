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

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	midleware.LogMiddleware(e)
	// routing with query parameter
	e.POST("/login", controller.LoginController)
	e.POST("/login/admin", controller.LoginAdminController)
	e.POST("/user", controller.CreateUserController)

	eJwt := e.Group("")
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

	bookmark := eJwt.Group("/bookmark")

	NewThreadControllers(eJwt)
	NewBookmarkedContoller(bookmark)
	Follow(eJwt)
	NewCommentControllers(eJwt)

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

func NewBookmarkedContoller(e *echo.Group) {
	e.GET("", controller.GetSaveThreadController)
	e.POST("", controller.CreateSaveThreadsController)
	e.DELETE("/:id", controller.DeleteSaveThreadsController)
}
func NewCommentControllers(e *echo.Group) {
	e.POST("/comment", controller.CreateCommentController)
	e.DELETE("/comment/:id", controller.DeleteCommentsControllerUser)
	e.PUT("/comment/:id", controller.UpdateCommentsControllerUser)
}
func Follow(e *echo.Group) {
	e.POST("/follow", controller.CreateFollowController)
	e.DELETE("/follow/:id", controller.DeleteFollowsControllerUser)
}
