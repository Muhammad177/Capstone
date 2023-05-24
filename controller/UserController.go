package controller

import (
	"Capstone/database"
	"Capstone/midleware"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"Capstone/models"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	// Bind data pengguna dari permintaan
	user := models.User{}
	err := c.Bind(&user)
	user.Role = "user"

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Menerima file foto dari permintaan
	file, err := c.FormFile("photo")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to upload photo")
	}

	// Generate nama unik untuk file foto
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Buka file foto yang diunggah
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open photo")
	}
	defer src.Close()

	// Simpan file foto di direktori lokal
	dstPath := "uploads/" + filename
	dst, err := os.Create(dstPath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save photo")
	}
	defer dst.Close()

	// Salin isi file foto yang diunggah ke file tujuan
	if _, err = io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save photo")
	}

	// Set path file foto pengguna
	user.Photo = dstPath

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

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}

	token, err := midleware.CreateToken(int(user.ID), user.Username, "user")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}
	usersResponse := models.UserResponse{int(user.ID), user.Username, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    usersResponse,
	})

}
