package controller

import (
	"errors"
	"net/http"
	"thread/model"
	"thread/service"

	"github.com/labstack/echo/v4"
)

type threadControllerImpl struct {
	srv service.ThreadService
}

var (
	errInvalidType = errors.New("invalid data type")
	errDatabase    = errors.New("error database")
	errNotFound    = errors.New("not found")
	getSuccess     = string("success get thread")
	createdSuccess = string("success created thread")
	updateSuccess  = string("success update thread")
	deleteSuccess  = string("success delete thread")
)

// CreateItem implements ItemController
func (t *threadControllerImpl) CreateThread(c echo.Context) error {
	var thread model.Thread
	if err := c.Bind(&thread); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errInvalidType.Error(),
		})
	}

	if err := t.srv.CreateThread(&thread, c.Request().Context()); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errDatabase.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": createdSuccess,
	})
}

// DeleteItem implements ItemController
func (t *threadControllerImpl) DeleteThread(c echo.Context) error {
	var thread model.Thread
	thread.ID = c.Param("id")

	if err := t.srv.DeleteThread(&thread, c.Request().Context()); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errNotFound.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": deleteSuccess,
	})
}

// GetAllItems implements ItemController
func (t *threadControllerImpl) GetAllThreads(c echo.Context) error {
	NameKey := c.FormValue("keyword")

	threads, err := t.srv.GetAllThreads(NameKey, c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errNotFound.Error(),
			"data":    threads,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": getSuccess,
		"data":    threads,
	})
}

// GetItemById implements ItemController
func (t *threadControllerImpl) GetThreadById(c echo.Context) error {
	var thread model.Thread
	thread.ID = c.Param("id")

	err := t.srv.GetThreadById(&thread, c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errNotFound.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": getSuccess,
		"data":    thread,
	})
}

// UpdateItem implements ItemController
func (t *threadControllerImpl) UpdateThread(c echo.Context) error {
	var thread model.Thread
	thread.ID = c.Param("id")

	if err := c.Bind(&thread); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errInvalidType,
		})
	}

	if err := t.srv.UpdateThread(&thread, c.Request().Context()); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errDatabase,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": updateSuccess,
	})
}

func NewThreadController(srv service.ThreadService) ThreadController {
	return &threadControllerImpl{
		srv: srv,
	}
}
