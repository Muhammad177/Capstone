package controller

import (
	"Remidi/database"
	"Remidi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCategorysController(c echo.Context) error {
	var category []models.Category
	err := database.DB.Find(&category).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all Category",
		"Produk":   category,
	})
}

// create new Category
func CreateCategoryController(c echo.Context) error {
	category := models.Category{}
	c.Bind(&category)

	if err := database.DB.Save(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Category",
		"Produk":  category,
	})
}

// delete Category by id
func DeleteCategoryController(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := database.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete category by id",
		"Produk":  category,
	})
}

// update Category by id
func UpdateCategoryController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var category models.Category
	category.ID = uint(id)
	if err := database.DB.Model(&models.Product{}).Where("id = ?", id).First(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := c.Bind(&category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := database.DB.Model(&models.Product{}).Updates(category).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product updated successfully",
		"user":    category,
	})
}
