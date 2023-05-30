package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/pkg/log"
)

type Bind struct {
	v *validator.Validate
}

func NewBind(v *validator.Validate) *Bind {
	if err := v.RegisterValidation("name", validateName); err != nil {
		log.WithErr(err).Warn("can't validate name fields")
	}

	if err := v.RegisterValidation("email", validateEmail); err != nil {
		log.WithErr(err).Warn("can't validate email fields")
	}

	if err := v.RegisterValidation("uuid4", validateUUID4); err != nil {
		log.WithErr(err).Warn("can't validate uuid4 fields")
	}

	if err := v.RegisterValidation("enum", validateEnum); err != nil {
		log.WithErr(err).Warn("can't validate enums fields")
	}

	if err := v.RegisterValidation("password", validatePassword); err != nil {
		log.WithErr(err).Warn("can't validate password fields")
	}
	return &Bind{v: v}
}

func (b *Bind) Struct(s any) error {
	return b.v.Struct(s)
}

func (b *Bind) Var(s any, tag string) error {
	return b.v.Var(s, tag)
}

func HandleData[T any](handler func(*gin.Context, T) error, v *Bind) func(*gin.Context) error {
	return func(c *gin.Context) error {
		// TODO var validation
		var t T
		if err := c.ShouldBindJSON(&t); err != nil {
			return err
		} else if err = v.Struct(&t); err != nil {
			return err
		}

		return handler(c, t)

	}
}
