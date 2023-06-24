package controller

import (
	"net/http"
	"strconv"

	"Capstone/database"
	"Capstone/dto"
	"Capstone/midleware"
	"Capstone/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateReportController(c echo.Context) error {
	Report := models.Report{}
	c.Bind(&Report)
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	Report.UserID = uint(id)
	newReport, err := database.CreateReport(c.Request().Context(), Report)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Report thread",
		"data":    newReport,
	})
}

func DeleteReportController(c echo.Context) error {
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var users models.User
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteReports(c.Request().Context(), Id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting Report data",
	})
}
func GetReportsController(c echo.Context) error {
	Report, err := database.GetReports(c.Request().Context())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Reports",
		"data":    dto.NewGetReportsResponse(Report),
	})
}
func GetReportByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	Report, err := database.GetReportByID(c.Request().Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Report",
		"data":    dto.NewGetReportResponse(Report),
	})
}
