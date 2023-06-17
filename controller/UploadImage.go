package controller

import (
	"Capstone/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Add your Cloudinary product environment credentials.
func UploadImageController(c echo.Context) error {

	file, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	resp, err := database.UploadImageCloud(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "image upload success",
		"Data":    resp,
	})
}
