package routes

import (
	"Capstone/constant"
	"Capstone/controller"
	"Capstone/dto"
	"Capstone/midleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(e *echo.Echo) {
	e.Validator = dto.NewValidator(validator.New())
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
	e.POST("/uploadImage", controller.UploadImageController)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.PUT("/admin/:id", controller.UpdateUserAdminController)
	eJwt.PUT("/admin", controller.UpdateDataAdminController)
	eJwt.DELETE("/admin/:id", controller.DeleteUserAdminController)
	eJwt.GET("/admin", controller.GetUsersAdminController)
	eJwt.GET("/admin/:id", controller.GetUserByidAdminController)
	eJwt.PUT("/user", controller.UpdateUserController)
	eJwt.DELETE("/user", controller.DeleteUserController)
	eJwt.GET("/user", controller.GetUserController)
	eJwt.GET("/Alluser", controller.GetAllUserController)
	//confirm

	bookmark := eJwt.Group("/bookmark")

	NewThreadControllers(eJwt)
	NewBookmarkedContoller(bookmark)
	Follow(eJwt)
	Like(eJwt)
	NewReportController(eJwt)
	NewCommentControllers(eJwt)
	MuteBlock(eJwt)
}

func NewThreadControllers(e *echo.Group) {
	e.GET("/admin/threads", controller.GetThreadController)
	e.GET("/threads/:id", controller.GetThreadsIDController)
	e.POST("/threads", controller.CreateThreadsController)
	e.DELETE("/admin/threads/:id", controller.DeleteThreadsControllerAdmin)
	e.DELETE("/threads/:id", controller.DeleteCommentsControllerUser)
	e.PUT("/admin/threads/:id", controller.UpdateThreadsControllerAdmin)
	e.PUT("/threads/:id", controller.UpdateThreadsControllerUser)
	e.GET("/threads", controller.GetThreadControllerByTitle)
	e.GET("/Allthreads", controller.GetAllThreadUserController)

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
	e.GET("/comment/:id", controller.GetCommentIDController)
	e.GET("/thread/:id/comment", controller.GetCommentController)
}
func Follow(e *echo.Group) {
	e.POST("/follow", controller.CreateFollowController)
	e.DELETE("/follow/:id", controller.DeleteFollowsControllerUser)
	e.GET("/follow", controller.GetFollowIDController)
}
func Like(e *echo.Group) {
	e.POST("/like", controller.CreateLikeController)
	e.DELETE("/like/:id", controller.DeleteLikeController)
	e.GET("/like", controller.GetLikeController)
}
func NewReportController(e *echo.Group) {
	e.POST("/report", controller.CreateReportController)
	e.DELETE("/report/:id", controller.DeleteReportController)
	e.GET("/report", controller.GetReportsController)
	e.GET("/report/:id", controller.GetReportByIdController)
}

func MuteBlock(e *echo.Group) {
	e.POST("/Mute", controller.CreateMuteController)
	e.DELETE("/Mute/:id", controller.DeleteMutesControllerUser)
	e.POST("/Block", controller.CreateBlockController)
	e.DELETE("/Block/:id", controller.DeleteBlockControllerUser)
}
