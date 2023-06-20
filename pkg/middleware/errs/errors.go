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

// 400
const (
	invalidValidation    ErrCode = "invalid_validation"
	codeInvalidOrExpired ErrCode = "code_invalid_or_expired"
	invalidPassword      ErrCode = "invalid_password"
	notFound             ErrCode = "not_found"
	alreadyExist         ErrCode = "already_exist"
	passwordNotSet       ErrCode = "password_not_set"
	websocketExcepted    ErrCode = "websocket_excepted"
)

// 500
const (
	cantSendMail ErrCode = "cant_send_mail"
	txFailed     ErrCode = "transaction_failed"
	serverError  ErrCode = "server_error"
)

type AbstractError struct {
	Status      int               `json:"-"`
	Code        ErrCode           `json:"code,omitempty"`
	Description string            `json:"description,omitempty"`
	Fields      map[string]string `json:"fields,omitempty"`
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

		switch err.(type) {
		case MyError:
			my = err.(MyError).GetInfo()

		case validator.ValidationErrors:
			my = newValidError(err.(validator.ValidationErrors))

		case *ent.NotFoundError:
			my = EntNotFoundError.AddError(err.(*ent.NotFoundError)).GetInfo()

		case *ent.ValidationError:
			valid := err.(*ent.ValidationError)
			my = EntValidError.AddError(valid).SetFields(
				map[string]string{valid.Name: fmt.Sprintf("%s is incorrect", valid.Name)}).GetInfo()

		case *ent.ConstraintError:
			my = EntConstraintError.AddError(err).GetInfo()

		case redis.Error:
			redisErr := err.(redis.Error)

			if redisErr == redis.Nil {
				my = RedisNilError.AddError(err).GetInfo()

			} else if redisErr == redis.TxFailedErr {
				my = RedisTxError.AddError(err).GetInfo()

			} else {
				my = RedisError.AddError(err).GetInfo()
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
