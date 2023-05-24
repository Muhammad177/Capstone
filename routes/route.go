package routes

import (
	"Remidi/constant"
	"Remidi/controller"
	"Remidi/midleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	midleware.LogMiddleware(e)
	// routing with query parameter
	e.POST("/login", controller.LoginUserController)
	e.POST("/user", controller.CreateUserController)
	e.GET("/user", controller.GetAllUsersController)
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("/items", controller.GetProductsController)
	eJwt.POST("/items", controller.CreateProductController)
	eJwt.GET("/items/:id", controller.GetProductController)
	eJwt.DELETE("/items/:id", controller.DeleteProductController)
	eJwt.PUT("/items/:id", controller.UpdateProductController)
	eJwt.GET("/category/:category_id", controller.GetCategoryControllerById)
	eJwt.GET("/Items", controller.GetCategoryControllerByName)

	eJwt.GET("/category", controller.GetCategorysController)
	eJwt.POST("/category", controller.CreateCategoryController)
	eJwt.GET("/category/:id", controller.GetCategoryControllerById)
	eJwt.DELETE("/category/:id", controller.DeleteCategoryController)
	eJwt.PUT("category/:id", controller.UpdateCategoryController)
	e.Logger.Fatal(e.Start(":8000"))
	return e
}
