package dto

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator(validator *validator.Validate) *CustomValidator {
	validator.RegisterValidation("passwordString", validatePasswordString)
	return &CustomValidator{
		Validator: validator,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// Jika terjadi kesalahan validasi internal
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}

		validationErrors := err.(validator.ValidationErrors)
		errorMessage := ""
		for _, fieldError := range validationErrors {
			// Customize  pesan error menjadi pesan yang lebih informatif
			switch fieldError.Tag() {
			case "required":
				errorMessage += fmt.Sprintf(`Field %s is required`+", ", fieldError.Field())
			case "email":
				errorMessage += fmt.Sprintf("Field %s must be a valid email address"+", ", fieldError.Field())
			case "min":
				errorMessage += fmt.Sprintf("Field %s minimum %s character"+", ", fieldError.Field(), fieldError.Param())
			case "max":
				errorMessage += fmt.Sprintf("Field %s maximum %s character"+", ", fieldError.Field(), fieldError.Param()) //bingung message
			case "number":
				errorMessage += fmt.Sprintf("Field %s must be number"+", ", fieldError.Field())
			default:
				errorMessage += fmt.Sprintf("Field %s is invalid"+"\n", fieldError.Field())
			}
		}
		return errors.New(errorMessage)
	}

	return nil
}

func validatePasswordString(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	regex, _ := regexp.Compile("^[A-Za-z0-9!@#$%^&*()-=_+]+$")
	result := regex.MatchString(password)
	return result
}
