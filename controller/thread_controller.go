package controller

import "github.com/labstack/echo/v4"

type ThreadController interface {
	GetAllThreads(c echo.Context) error
	CreateThread(c echo.Context) error
	UpdateThread(c echo.Context) error
	DeleteThread(c echo.Context) error
	GetThreadById(c echo.Context) error
}
