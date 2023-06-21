package controller

import (
	"Capstone/database"
	"Capstone/midleware"
	"Capstone/models"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	// Retrieve the user ID from the JWT token
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Fetch the user's information based on the user ID
	var users []models.User
	if err := database.DB.Preload("Threads").Where("id = ?", int(id)).Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := make([]models.AllUserFollow, len(users))
	for i, users := range users {
		user[i] = models.ConvertUserToAllUserFollow(&users)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user info by id",
		"user":    user,
	})
}
func UpdateUserController(c echo.Context) error {
	// Retrieve the JWT token from the request context and extract the role claim
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var users models.User
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	if err := c.Validate(users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error update user",
			"error":    err.Error(),
		})
	}

	if err := database.DB.Model(&users).Updates(users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    users,
	})
}
func DeleteUserController(c echo.Context) error {
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var users models.User
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete User by id",
		"Produk":  users,
	})
}
func LoginController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}

	token, err := midleware.CreateToken(int(user.ID), user.Username, user.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}
	usersResponse := models.UserResponse{int(user.ID), user.Username, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Login",
		"user":    usersResponse,
	})

}

func GetAllUserController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "User" {
		return c.JSON(http.StatusUnauthorized, "Error Account")
	}

	var users []models.User
	err = database.DB.Find(&users).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve users from the database")
	}
	allUsers := make([]models.AllUserSearch, len(users))
	for i, user := range users {
		allUsers[i] = models.ConvertAllUserSearch(&user)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success: Retrieved all users",
		"users":   allUsers,
	})
}
func GetAllThreadUserController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "User" {
		return c.JSON(http.StatusUnauthorized, "Error Account")
	}
	thread, err := database.GetThreads(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	allThreads := make([]models.AllThread, len(thread))
	for i, thread := range thread {
		allThreads[i] = models.ConverThreadToAllThread(&thread)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success: Retrieved all threads",
		"data":    allThreads,
	})
}
