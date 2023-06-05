package controller

import (
	"Capstone/database"
	"Capstone/midleware"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"

	"Capstone/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	// Retrieve the JWT token from the request context and extract the role claim
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := claims["user_id"].(float64)
	var users []models.User
	if err := database.DB.Where("id = ?", id).Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user info by id",
		"user":    users,
	})
}
func UpdateUserController(c echo.Context) error {
	// Retrieve the JWT token from the request context and extract the role claim
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(float64)

	var users models.User
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(float64)
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
func GetImageHandler(c echo.Context) error {
	
	// Dapatkan UUID gambar dari parameter permintaan
	uuid := c.Param("uuid")

	// Construct the file path based on the UUID string
	filePath := filepath.Join("uploads", uuid) // Folder "uploads" berada dalam direktori saat ini

	// Dapatkan tipe MIME file
	file, err := os.Open(filePath)
	if err != nil {
		return c.String(http.StatusNotFound, "File not found")
	}
	defer file.Close()

	// Baca awal file untuk mendapatkan tipe MIME
	buffer := make([]byte, 512) // Membaca 512 byte pertama
	_, err = file.Read(buffer)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	// Deteksi tipe MIME
	mimeType := mime.TypeByExtension(filepath.Ext(filePath))
	if mimeType == "" {
		mimeType = http.DetectContentType(buffer)
	}

	// Set header Content-Type pada response
	c.Response().Header().Set("Content-Type", mimeType)

	// Baca file dengan ekstensi yang tepat
	imageBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	// Kirim data gambar sebagai response
	return c.Blob(http.StatusOK, mimeType, imageBytes)
}
