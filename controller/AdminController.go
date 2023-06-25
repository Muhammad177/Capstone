package controller

import (
	"Capstone/database"
	"Capstone/midleware"
	"Capstone/models"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	// Bind data pengguna dari permintaan
	user := models.User{}
	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create user",
			"error":    err.Error(),
		})
	}

	if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		// Email sudah ada, kembalikan respons error
		return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
	}
	// Set nilai default untuk role
	user.Role = "User"
	// Simpan pengguna ke database
	err = database.DB.Save(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save user")
	}

	// Mengembalikan respons JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil membuat pengguna baru",
		"user":    user,
	})
}
func UpdateUserAdminController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}

	id := c.Param("id")

	var users models.User
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	previousEmail := users.Email

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	if err := c.Validate(users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error Update user",
			"error":    err.Error(),
		})
	}
	if previousEmail != users.Email {
		var existingUser models.User
		if err := database.DB.Where("email = ?", users.Email).First(&existingUser).Error; err == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
		}
	}
	if err := database.DB.Model(&users).Updates(users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    users,
	})
}
func UpdateDataAdminController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}

	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var users models.User
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	previousEmail := users.Email

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	if err := c.Validate(users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error Update user",
			"error":    err.Error(),
		})
	}
	if previousEmail != users.Email {
		var existingUser models.User
		if err := database.DB.Where("email = ?", users.Email).First(&existingUser).Error; err == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
		}
	}
	if err := database.DB.Model(&users).Updates(users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    users,
	})
}

func GetUserByidAdminController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}

	id := c.Param("id")
	var users models.User
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get User by id",
		"user":    users,
	})
}
func GetUsersAdminController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}

	var users []models.User
	err = database.DB.Find(&users).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve users from the database")
	}
	allUsers := make([]models.AllUser, len(users))
	for i, user := range users {
		allUsers[i] = models.ConvertUserToAllUser(&user)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success: Retrieved all users",
		"users":   allUsers,
	})
}
func DeleteUserAdminController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}

	id := c.Param("id")
	var users models.User

	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user by ID",
		"user":    users,
	})
}

func LoginAdminController(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	// Verifikasi admin berdasarkan email dan password
	if user.Email == "admin@gmail.com" && user.Password == "admin123" {
		token, err := midleware.CreateToken(int(user.ID), user.Username, "admin")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed Login",
				"error":   err.Error(),
			})
		}

		adminResponse := models.UserResponse{ID: int(user.ID), Name: user.Username, Email: user.Email, Token: token}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login Admin Sukses",
			"Admin":   adminResponse,
		})
	}

	if err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}
	if user.Role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}
	token, err := midleware.CreateToken(int(user.ID), user.Username, user.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}
	usersResponse := models.UserResponse{ID: int(user.ID), Name: user.Username, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Login Admin",
		"user":    usersResponse,
	})
}
