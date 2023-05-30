package errs

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"net/http"
)

// Sign-in errors
var (
	CodeError        = newStandardError(http.StatusBadRequest, "Code is not correct", "Try to request a new one")
	PasswordError    = newStandardError(http.StatusBadRequest, "Wrong password", "You can still sign in by your email!")
	PasswordNotFound = newStandardError(http.StatusBadRequest, "You have not registered a password for you account", "Try change the password in your profile")
)

// Auth errors
var (
	UnAuthorized = newStandardError(http.StatusUnauthorized, "You are not logged in", "Click on the button below to sign in!")
)

// Server errors
var (
	ServerError = newStandardError(http.StatusInternalServerError, "Server exception was occurred", "Try to restart the page")
	EmailError  = newStandardError(http.StatusInternalServerError, "Can't send message to your email", "Try to send it later")
)

type MyError interface {
	GetInfo() *AbstractError
}

type AbstractError struct {
	Status int `json:"-"`
	Msg    any
	Advice string `json:"advice,omitempty"`
	Err    error  `json:"-"`
}

type ErrHandler struct {
	log *log.Logger
}

func NewErrHandler() *ErrHandler {
	return &ErrHandler{log: log.NewLogger(log.ErrLevel, &log.JSONFormatter{}, true)}
}

// HandleErrors of StandardError type
func (e *ErrHandler) HandleErrors(handler func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err == nil {
			return
		}

		my := ServerError.GetInfo()

		if myErr, ok := err.(MyError); ok {
			my = myErr.GetInfo()

		} else if vErrs, ok := err.(validator.ValidationErrors); ok {
			my = newValidError(vErrs).GetInfo()

		} else if entErr, ok := err.(EntError); ok {
			//TODO generate EntError by ent hooks
			my = entErr.GetInfo()
		} else if redisErr, ok := err.(RedisError); ok {
			//TODO generate RedisError by redis hooks
			my = redisErr.GetInfo()
		}

		e.log.WithErr(err).Err(my.Msg)
		c.JSON(my.Status, my)
	}
}
