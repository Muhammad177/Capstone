package midleware

import (
	"Capstone/constant"
	"Capstone/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID int, name string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}
func ClaimsId(c echo.Context) (float64, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := claims["user_id"].(float64)
	return id, nil
}
func ClaimsRole(c echo.Context) (string, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	role := claims["role"].(string)
	return role, nil
}
func CheckMutekStatus(mutes []models.Mute, userID float64) (bool, string) {
	for _, m := range mutes {
		if m.UserID == userID {
			if m.Status == "mute" {
				return true, "You have been muted by an admin"
			} else if m.Status == "block" {
				return true, "You have been blocked by an admin"
			}
		}
	}
	return false, ""
}
func CheckBlockStatus(mutes []models.Mute, userID float64) (bool, string) {
	for _, m := range mutes {
		if m.UserID == userID {
			if m.Status == "block" {
				return true, "You have been blocked by an admin"
			}
		}
	}
	return false, ""
}
