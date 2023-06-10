package controller

import (
	"Capstone/database"
	"Capstone/midleware"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"strings"

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
	if err := database.DB.Preload("Threads").Preload("Comment").Where("id = ?", int(id)).Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user info by id",
		"user":    users,
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
func LogoutController(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")

	// Memastikan header Authorization tidak kosong dan memiliki format "Bearer <token>"
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid token",
		})
	}

	// Mengambil token saja (tanpa "Bearer ")
	token := strings.TrimPrefix(tokenString, "Bearer ")

	// Mendekode token untuk mendapatkan user_id
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Menghancurkan token
	err = midleware.DestroyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to destroy token",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logout successful",
		"user_id": id,
	})
}

func GetImageHandler(c echo.Context) error {
	// Dapatkan UUID gambar dari parameter permintaan
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var users models.User
	if err := database.DB.Select("photo").Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	photo := users.Photo

	// Construct the file path based on the UUID string
	filePath := filepath.Join("uploads", photo) // Folder "uploads" berada dalam direktori saat ini

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
