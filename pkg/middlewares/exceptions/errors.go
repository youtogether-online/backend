package exceptions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Sign-in errors
var (
	SignInUnknown = newError(http.StatusNotFound, "There is no such user", "But you can still find another existing user!")
	CodeError     = newError(http.StatusBadRequest, "Code is not correct", "Try to request a new one")
	PasswordError = newError(http.StatusBadRequest, "Wrong password", "You can still sign in by your email!")
)

// Auth errors
var (
	PermissionError = newError(http.StatusForbidden, "You don't have this permission", "Try to ask the owner about it")
	UnAuthorized    = newError(http.StatusUnauthorized, "You are not logged in", "Click on the button below to sign in!")
)

// Input errors
var (
	ValidError = newError(http.StatusBadRequest, "Validation error", "Try to enter the correct data")
	DataError  = newError(http.StatusBadRequest, "Insufficient data", "Try to enter the remaining data")
)

// ServerError err
var (
	ServerError = newError(http.StatusInternalServerError, "Server exception was occurred", "Try to restart the page")
)

// ErrorHandler used for error handling. Handles only MyError type errors
func ErrorHandler(c *gin.Context) {
	c.Next()

	errs := c.Errors

	if errs.Last() == nil {
		return
	}

	for i, err := range errs {
		if my, ok := err.Err.(MyError); ok {

			logrus.WithError(my.Err).WithField("index", i).Error(my.Msg)
			res := gin.H{"error": my.Msg, "advice": my.Advice}

			if vErrs, ok := my.Err.(validator.ValidationErrors); ok {
				fields := make(gin.H)

				for _, vErr := range vErrs {
					field := vErr.Field()
					switch vErr.Tag() {
					case "email":
						fields[field] = fmt.Sprintf("%s is not an email", field)
					case "required":
						fields[field] = fmt.Sprintf("%s should not be empty", field)
					case "numeric":
						fields[field] = fmt.Sprintf("%s must be a number", field)
					case "gte":
						fields[field] = fmt.Sprintf("%s must be greater or equal %s", field, vErr.Param())
					case "lte":
						fields[field] = fmt.Sprintf("%s must be lesser or equal %s", field, vErr.Param())
					case "len":
						fields[field] = fmt.Sprintf("%s must have a length of %s", field, vErr.Param())
					case "gt":
						fields[field] = fmt.Sprintf("%s must be greater than %s", field, vErr.Param())
					case "lt":
						fields[field] = fmt.Sprintf("%s must be lesser than %s", field, vErr.Param())
					}
				}
				res["fields"] = fields
			}

			if i == 0 {
				c.JSON(my.Code, res)
			}
		} else {
			logrus.WithError(err.Err).Error("UNEXPECTED ERROR")
			if i == 0 {
				c.JSON(ServerError.Code, ServerError.Msg)
			}
		}
	}
}
