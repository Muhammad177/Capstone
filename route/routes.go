package route

import (
	"thread/controller"
	"thread/repository"
	"thread/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB) *echo.Echo {
	e := echo.New()

	repo := repository.NewThreadRepository(db)
	threadSv := service.NewThreadService(repo)
	threadCtr := controller.NewThreadController(threadSv)

	e.GET("/threads", threadCtr.GetAllThreads)
	e.GET("/threads/:id", threadCtr.GetThreadById)
	e.PUT("/threads/:id", threadCtr.UpdateThread)
	e.DELETE("/threads/:id", threadCtr.DeleteThread)
	e.POST("/threads", threadCtr.CreateThread)

	return e
}
