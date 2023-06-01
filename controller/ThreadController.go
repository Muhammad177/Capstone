package controller

import (
	"net/http"
	"strconv"

	"Capstone/database"
	"Capstone/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetThreadController(c echo.Context) error {
	thread, err := database.GetThreads(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Thread",
		"data":    thread,
	})

}

func GetThreadsIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	thread, err := database.GetThreadsByID(c.Request().Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Thread",
		"data":    thread,
	})
}

func CreateThreadsController(c echo.Context) error {
	thread := models.Thread{}
	c.Bind(&thread)

	newThread, err := database.CreateThreads(c.Request().Context(), thread)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating thread",
		"data":    newThread,
	})
}

func DeleteThreadsController(c echo.Context) error {
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
		"message": "success deleting thread data",
	})
}

func UpdateThreadsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	thread := models.Thread{}
	c.Bind(&thread)

	updateThread, err := database.UpdateThreads(c.Request().Context(), id, thread)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating thread data",
		"data":    updateThread,
	})
}
