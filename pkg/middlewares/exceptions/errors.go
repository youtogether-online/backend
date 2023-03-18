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
	LoginUnknown  = newError(http.StatusNotFound, "User does not exist")
	CodeError     = newError(http.StatusBadRequest, "Code is not correct")
	PasswordError = newError(http.StatusBadRequest, "Wrong password")
)

// Auth errors
var (
	PermissionError = newError(http.StatusForbidden, "You don't have this permission")
	UnAuthorized    = newError(http.StatusUnauthorized, "You are not logged in")
)

// Input errors
var (
	ValidError = newError(http.StatusBadRequest, "Validation error")
	DataError  = newError(http.StatusBadRequest, "Insufficient data")
)

// ServerError err
var (
	ServerError = newError(http.StatusInternalServerError, "Server exception was occurred")
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
			res := gin.H{"Error": my.Msg}

			if vErrs, ok := my.Err.(validator.ValidationErrors); ok {
				fields := make(gin.H)

				for _, vErr := range vErrs {
					switch vErr.Tag() {
					case "email":
						fields[vErr.Field()] = fmt.Sprintf("%s is incorrect", vErr.Field())
					case "required":
						fields[vErr.Field()] = fmt.Sprintf("%s should not be empty", vErr.Field())
					case "numeric":
						fields[vErr.Field()] = fmt.Sprintf("%s must be a number", vErr.Field())
					case "gte":
						fields[vErr.Field()] = fmt.Sprintf("%s must be greater or equal %s", vErr.Field(), vErr.Param())
					case "lte":
						fields[vErr.Field()] = fmt.Sprintf("%s must be lesser or equal %s", vErr.Field(), vErr.Param())
					case "len":
						fields[vErr.Field()] = fmt.Sprintf("%s must have a length of %s", vErr.Field(), vErr.Param())
					case "gt":
						fields[vErr.Field()] = fmt.Sprintf("%s must be greater than %s", vErr.Field(), vErr.Param())
					case "lt":
						fields[vErr.Field()] = fmt.Sprintf("%s must be lesser than %s", vErr.Field(), vErr.Param())
					}
				}
				res["Fields"] = fields
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
