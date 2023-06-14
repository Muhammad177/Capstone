package controller

import (
	"net/http"
	"strconv"

	"Capstone/database"
	"Capstone/models"

	"github.com/labstack/echo/v4"
)

func CreateLikeThreads(c echo.Context) error {
	user := c.Get("user").(models.User)
	thread_id, _ := strconv.Atoi(c.Param("id"))

	_, err := database.CreateLikeThreads(c.Request().Context(), int(user.ID), thread_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Liked")
}

func DeleteLikeThreads(c echo.Context) error {
	user := c.Get("user").(models.User)
	thread_id, _ := strconv.Atoi(c.Param("id"))

	err := database.DeleteLikeThreads(c.Request().Context(), int(user.ID), thread_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Unliked")
}

func GetLikeThreads(c echo.Context) error {
	user := c.Get("user").(models.User)

	threads, err := database.GetLikeThreads(c.Request().Context(), int(user.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, threads)
}
