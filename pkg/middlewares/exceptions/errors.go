package exceptions

import (
	"github.com/gin-gonic/gin"
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

	if c.Errors.Last() == nil {
		return
	}

	for i, err := range c.Errors {
		if my, ok := err.Err.(MyError); ok {
			logrus.WithField("err", my.Err).Errorf("%v: %s", i, my.Msg)
			if i == 0 {
				c.JSON(my.Code, gin.H{"error": my.Msg})
			}
		}
	}
}
