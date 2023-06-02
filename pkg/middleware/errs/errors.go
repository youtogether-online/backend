package errs

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"net/http"
)

// Sign-in errors
var (
	CodeError        = newStandardError(http.StatusBadRequest, "Code is not correct", "Try to request a new one")
	PasswordError    = newStandardError(http.StatusBadRequest, "Wrong password", "You can still sign in by your email!")
	PasswordNotFound = newStandardError(http.StatusBadRequest, "You have not registered a password for you account", "Try change the password in your profile")
)

// ent error templates
var (
	EntNotFoundError    = newEntError(http.StatusBadRequest, "Entity is not found", "Try to find another entity")
	EntValidError       = newEntError(http.StatusBadRequest, "", "")
	EntConstraintError  = newEntError(http.StatusBadRequest, "Can't set this value", "Try to get another existing value")
	EntNotSingularError = newEntError(http.StatusInternalServerError, "An object was expected, but several were found", "Try to look for something else")
	EntNotLoadedError   = newEntError(http.StatusInternalServerError, "Can't load data", "Try to request it later")

	RedisNilError = NewRedisError(http.StatusBadRequest, "Can't find value", "Maybe this value is not registered?")
	RedisTxError  = NewRedisError(http.StatusInternalServerError, "Operation failed", "Try to request it later")
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
	Status int               `json:"-"`
	Msg    string            `json:"message,omitempty"`
	Fields map[string]string `json:"fields,omitempty"`
	Advice string            `json:"advice,omitempty"`
	Err    error             `json:"-"`
}

type ErrHandler struct {
	log *log.Logger
}

func NewErrHandler() *ErrHandler {
	return &ErrHandler{log: log.NewLogger(log.ErrLevel, &log.JSONFormatter{}, true)}
}

func (e *ErrHandler) HandleError(handler func(*gin.Context) error) gin.HandlerFunc {
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
			my = entErr.GetInfo()

		} else if redisErr, ok := err.(redis.Error); ok {
			switch redisErr {
			case redis.Nil:
				my = RedisNilError.AddError(err).GetInfo()
			default:
				my = RedisTxError.AddError(err).GetInfo()
			}
		}

		entry := e.log.WithErr(err)

		if my.Fields == nil {
			entry.Err(my.Msg)
		} else {
			entry.Err(my.Fields)
		}

		c.JSON(my.Status, my)
	}
}
