package errs

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"net/http"
)

type MyError interface {
	GetInfo() *AbstractError
}

type ErrCode string

// 400
const (
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

		entry := e.log.WithErr(err)
		my := &AbstractError{
			Status:      http.StatusInternalServerError,
			Code:        serverError,
			Description: "Server exception was occurred",
		}

		switch err.(type) {
		case StandardError:
			stErr := err.(StandardError)
			entry.Err(stErr.description)
			my = stErr.GetInfo()

		case validator.ValidationErrors:
			my = newValidError(err.(validator.ValidationErrors))
			entry.Err(my.Fields)

		case *ent.ValidationError:
			my = newValidErrorEnt(err.(*ent.ValidationError))
			entry.Err(my.Fields)

		case *ent.NotFoundError:
			my = EntNotFoundError.GetInfo()
			entry.Err(my.Description)

		case *ent.ConstraintError:
			my = EntConstraintError.GetInfo()
			entry.Err(my.Description)

		case redis.Error:
			redisErr := err.(redis.Error)

			if redisErr == redis.Nil {
				my = RedisNilError.GetInfo()

			} else if redisErr == redis.TxFailedErr {
				my = RedisTxError.GetInfo()

			} else {
				my = RedisError.GetInfo()
			}
			entry.Err(my.Description)
		}

		c.JSON(my.Status, my)
	}
}
