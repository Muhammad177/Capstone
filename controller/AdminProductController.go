package controller

import (
	"Remidi/database"
	"Remidi/midleware"

	"Remidi/models"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

var Products []models.Product

// get all Products
func GetProductsController(c echo.Context) error {

	err := database.DB.Preload("Category").Find(&Products).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all Products",
		"Produk":   Products,
	})
}
func GetAllUsersController(c echo.Context) error {
	var user []models.User
	err := database.DB.Find(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all User",
		"Produk":   user,
	})
}

// get Product by id
func GetProductController(c echo.Context) error {
	id := c.Param("id")
	var Products models.Product
	if err := database.DB.Preload("Category").Where("id = ?", id).First(&Products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Product by id",
		"user":    Products,
	})
}

// get Product by category_id
func GetCategoryControllerById(c echo.Context) error {
	category_id := c.Param("id")
	var product models.Product
	if err := database.DB.Preload("Category").Where("category_id = ?", category_id).Find(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Product by CategoryID",
		"user":    Products,
	})
}

func GetCategoryControllerByName(c echo.Context) error {
	keyword := c.QueryParam("keywoards")
	keyword = "%" + keyword + "%"
	var items []models.Product
	if err := database.DB.Preload("Category").Where("nama LIKE ?", keyword).Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully retrieved items by keyword",
		"items":   items,
	})
}

// delete Product by id
func DeleteProductController(c echo.Context) error {
	id := c.Param("id")
	var Products models.Product
	if err := database.DB.Where("id = ?", id).First(&Products).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&Products).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Product by id",
		"Produk":  Products,
	})
}

// update Product by id
func UpdateProductController(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := database.DB.Model(&models.Product{}).Where("id = ?", id).First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := database.DB.Model(&models.Product{}).Updates(product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product updated successfully",
		"user":    product,
	})
}

// create new Product
func CreateProductController(c echo.Context) error {
	Products := models.Product{}
	c.Bind(&Products)

	if err := database.DB.Save(&Products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := database.DB.Preload("Category").First(&Products).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch Product")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Product",
		"Produk":  Products,
	})
}
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}

	token, err := midleware.CreateToken(int(user.ID), user.Name, "user")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}
	usersResponse := models.UserResponse{int(user.ID), user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    usersResponse,
	})

}
