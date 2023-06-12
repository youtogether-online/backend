package errs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/pkg/log"
)

type MyError interface {
	GetInfo() *AbstractError
}

type ErrCode string

const (
	validErr    ErrCode = "validation"
	authErr     ErrCode = "authorization"
	serverErr   ErrCode = "server"
	notFoundErr ErrCode = "not_found"
)

type AbstractError struct {
	Status      int               `json:"-"`
	Code        ErrCode           `json:"type,omitempty"`
	Description string            `json:"description,omitempty"`
	Fields      map[string]string `json:"Fields,omitempty"`
	Advice      string            `json:"Advice,omitempty"`
	Err         error             `json:"-"`
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

		} else if notFound, ok := err.(*ent.NotFoundError); ok {
			my = EntNotFoundError.AddError(notFound).GetInfo()

		} else if valid, ok := err.(*ent.ValidationError); ok {
			my = EntValidError.AddError(valid).SetFields(
				map[string]string{valid.Name: fmt.Sprintf("%s is incorrect", valid.Name)}).GetInfo()

		} else if notSingular, ok := err.(*ent.NotSingularError); ok {
			my = EntNotSingularError.AddError(notSingular).GetInfo()

		} else if constraint, ok := err.(*ent.ConstraintError); ok {
			my = EntConstraintError.AddError(constraint).GetInfo()

		} else if notLoaded, ok := err.(*ent.NotLoadedError); ok {
			my = EntNotLoadedError.AddError(notLoaded).GetInfo()

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
			entry.Err(my.Description)
		} else {
			entry.Err(my.Fields)
		}

		c.JSON(my.Status, my)
	}
}
