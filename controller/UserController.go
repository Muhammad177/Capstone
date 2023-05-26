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

func CreatePhoto(c echo.Context) (string, error) {
	// Check if a file photo is present in the request
	_, err := c.FormFile("photo")
	if err != nil {
		return "", nil
	}

	// Menerima file foto dari permintaan
	file, err := c.FormFile("photo")
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Failed to upload photo")
	}

	// Generate nama unik untuk file foto
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Buka file foto yang diunggah
	src, err := file.Open()
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to open photo")
	}
	defer src.Close()

	// Simpan file foto di direktori lokal
	dstPath := "uploads/" + filename
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to save photo")
	}
	defer dst.Close()

	// Salin isi file foto yang diunggah ke file tujuan
	if _, err = io.Copy(dst, src); err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to save photo")
	}

	// Mengembalikan path file foto
	return dstPath, nil
}

func CreateUserController(c echo.Context) error {
	// Bind data pengguna dari permintaan
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		// Email sudah ada, kembalikan respons error
		return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
	}
	// Set nilai default untuk role
	user.Role = "User"

	// Simpan foto pengguna
	photoPath, err := CreatePhoto(c)
	if err != nil {
		return err
	}

	// Set path file foto pengguna
	user.Photo = photoPath

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
func UpdateUserController(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	previousEmail := user.Email

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if previousEmail != user.Email {
		var existingUser models.User
		if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
		}
	}

	// Check if new photo is uploaded
	_, err := c.FormFile("photo")
	if err == nil {
		// New photo is uploaded, execute CreatePhoto function
		photoPath, err := CreatePhoto(c)
		if err != nil {
			return err
		}

		// Delete previous photo
		if user.Photo != "" {
			if err := os.Remove(user.Photo); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user photo")
			}
		}

		user.Photo = photoPath
	} else if err == http.ErrMissingFile {
		// No new photo provided, check if existing photo needs to be deleted
		if user.Photo != "" {
			// Delete previous photo from database and local directory
			if err := os.Remove(user.Photo); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user photo")
			}

			user.Photo = ""
		}
	}

	// Update the user in the database
	if err := database.DB.Model(&user).Updates(map[string]interface{}{
		"photo": user.Photo,
	}).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    user,
	})
}

func GetUsersController(c echo.Context) error {
	var users []models.User
	err := database.DB.Find(&users).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}
func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Menghapus file foto terkait jika ada
	if user.Photo != "" {
		if err := os.Remove(user.Photo); err != nil {
			// Jika gagal menghapus file, Anda dapat menangani kesalahan di sini
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user photo")
		}
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user by ID",
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
