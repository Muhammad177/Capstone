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

func CreateCommentController(c echo.Context) error {
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mutes, err := database.GetMute(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	isMutedOrBlocked, message := midleware.CheckMutekStatus(mutes, id)
	if isMutedOrBlocked {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": message,
		})
	}
	Comment := models.Comment{}
	c.Bind(&Comment)
	Comment.UserID = int(id)
	newComment, err := database.CreateComment(c.Request().Context(), Comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success creating comment",
		"data":    models.ConvertCommentToCommentResponse(&newComment),
	})
}
func DeleteCommentsControllerUser(c echo.Context) error {
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteComments(c.Request().Context(), commentID, int(id))
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting Comment data",
	})
}

func UpdateCommentsControllerUser(c echo.Context) error {
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	CId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	Comment := models.Comment{}
	c.Bind(&Comment)
	updateComment, err := database.UpdateComments(c.Request().Context(), int(id), CId, Comment)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating Comment data",
		"data":    updateComment,
	})
}
func GetCommentController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	comment, err := database.GetComments(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success: Retrieved all Comment",
		"data":    dto.NewGetCommentsResponse(comment),
	})
}

func GetCommentIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	comment, err := database.GetCommentID(c.Request().Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Comment",
		"data":    dto.NewGetCommentResponse(comment),
	})
}
