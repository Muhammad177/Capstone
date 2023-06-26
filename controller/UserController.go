package controller

import (
	"Capstone/database"
	"Capstone/dto"
	"Capstone/midleware"
	"Capstone/models"
	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	// Retrieve the user ID from the JWT token
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, followingCount, followerCount, err := database.GetUsersByID(c.Request().Context(), int(id))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user info by id",
		"data":    dto.NewDetailUserResponse(user, followingCount, followerCount),
	})
}
func UpdateUserController(c echo.Context) error {
	// Retrieve the JWT token from the request context and extract the role claim
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.User{}
	c.Bind(&user)

	err = database.UpdateUser(c.Request().Context(), int(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error update user",
			"error":    err.Error(),
		})
	}
	user.Role = "User"

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"data":    dto.ConvertUserToAllUser(user),
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteUser(c.Request().Context(), int(id))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete User by id",
	})
}

func LoginController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	savedUser, err := database.GetUserByEmail(c.Request().Context(), user.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}
	if user.Password != savedUser.Password {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Password incorect",
			"error":   err.Error(),
		})
	}

	token, err := midleware.CreateToken(int(savedUser.ID), savedUser.Username, savedUser.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}
	usersResponse := models.UserResponse{ID: int(savedUser.ID), Name: savedUser.Username, Email: savedUser.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Login",
		"user":    usersResponse,
	})

}

func GetAllUserController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "User" {
		return c.JSON(http.StatusUnauthorized, "Error Account")
	}

	var users []models.User
	err = database.DB.Find(&users).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve users from the database")
	}
	allUsers := make([]models.AllUserSearch, len(users))
	for i, user := range users {
		allUsers[i] = models.ConvertAllUserSearch(&user)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success: Retrieved all users",
		"users":   allUsers,
	})
}
func GetAllThreadUserController(c echo.Context) error {
	role, err := midleware.ClaimsRole(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if role != "User" {
		return c.JSON(http.StatusUnauthorized, "Error Account")
	}
	thread, err := database.GetThreads(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	allThreads := make([]models.AllThread, len(thread))
	for i, thread := range thread {
		allThreads[i] = models.ConverThreadToAllThread(&thread)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success: Retrieved all threads",
		"data":    dto.NewGetAllThreadResponses(thread),
	})
}

func FollowUserController(c echo.Context) error {
	followUser := dto.FollowUserRequest{}
	if err := c.Bind(&followUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if followUser.ID == int(id) {
		return echo.NewHTTPError(http.StatusBadRequest, ("Can't follow ur self"))
	}
	fmt.Println(followUser, id)

	err = database.FollowUser(c.Request().Context(), int(id), followUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success following user",
	})
}

func UnFollowUserController(c echo.Context) error {
	followUser := dto.FollowUserRequest{}
	if err := c.Bind(&followUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := midleware.ClaimsId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.UnFollowUser(c.Request().Context(), int(id), followUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success following user",
	})

}
