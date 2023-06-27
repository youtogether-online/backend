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

		my := &AbstractError{
			Status:      http.StatusInternalServerError,
			Code:        serverError,
			Description: "Server exception was occurred",
			Err:         err,
		}

		switch err.(type) {
		case StandardError:
			my = err.(StandardError).GetInfo()

		case validator.ValidationErrors:
			my = newValidError(err.(validator.ValidationErrors))

		case *ent.ValidationError:
			my = newValidErrorEnt(err.(*ent.ValidationError))

		case *ent.NotFoundError:
			my = EntNotFoundError.GetInfo(err.(*ent.NotFoundError))

		case *ent.ConstraintError:
			my = EntConstraintError.GetInfo(err.(*ent.ConstraintError))

		case redis.Error:
			redisErr := err.(redis.Error)
			switch err.(redis.Error) {
			case redis.Nil:
				my = RedisNilError.GetInfo(redisErr)

			case redis.TxFailedErr:
				my = RedisTxError.GetInfo(redisErr)

			default:
				my = RedisError.GetInfo(redisErr)
			}
		}

		l := e.log.WithErr(my.Err)
		if my.Fields == nil {
			l.Err(my.Description)
		} else {
			l.Err(my.Fields)
		}

		c.JSON(my.Status, my)
	}
}
