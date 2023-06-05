package controller

import (
	"net/http"
	"strconv"

	"Capstone/database"
	"Capstone/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetSavedthreadController(c echo.Context) error {
	savedthread, err := database.GetSavedthreads(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting save thread",
		"data":    savedthread,
	})

}

func GetSavedthreadsIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	savedthread, err := database.GetSavedthreadsByID(c.Request().Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting save thread by ID",
		"data":    savedthread,
	})
}

func CreateSavedthreadsController(c echo.Context) error {
	savedthread := models.Savedthread{}
	c.Bind(&savedthread)

	newSavedthread, err := database.CreateSavedthreads(c.Request().Context(), savedthread)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create saving thread",
		"data":    newSavedthread,
	})
}

func DeleteSavedthreadsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteThreads(c.Request().Context(), id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete savethread",
	})
}

func UpdateSavedthreadsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	savedthread := models.Savedthread{}
	c.Bind(&savedthread)

	updateSavedthread, err := database.UpdateSavedthreads(c.Request().Context(), id, savedthread)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update savethread",
		"data":    updateSavedthread,
	})
}
