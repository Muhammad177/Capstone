package controller

import (
	"Capstone/database"
	"Capstone/midleware"
	"Capstone/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateMuteController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}
	Mute := models.Mute{}
	c.Bind(&Mute)
	Mute.Status = "mute"
	newMute, err := database.CreateMute(c.Request().Context(), Mute)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Mute User",
		"data":    newMute,
	})
}

func DeleteMutesControllerUser(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteMutes(c.Request().Context(), Id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting Mute data",
	})
}
func CreateBlockController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}
	Mute := models.Mute{}
	c.Bind(&Mute)
	Mute.Status = "block"
	newMute, err := database.CreateMute(c.Request().Context(), Mute)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Mute User",
		"data":    newMute,
	})
}

func DeleteBlockControllerUser(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, "Only admin can access")
	}
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteMutes(c.Request().Context(), Id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting Block data",
	})
}
