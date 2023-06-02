package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/pkg/log"
)

func NewValidator() *validator.Validate {
	v := validator.New()
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
	return v
}

func HandleBody[T any](handler func(*gin.Context, T) error, v *validator.Validate) func(*gin.Context) error {
	return func(c *gin.Context) error {

		var t T
		if err := c.ShouldBindJSON(&t); err != nil {
			return err
		} else if err = v.Struct(&t); err != nil {
			return err
		}

		return handler(c, t)

	}
}

func HandleParam(handler func(*gin.Context, string) error, name string, tag string, v *validator.Validate) func(*gin.Context) error {
	return func(c *gin.Context) error {
		t := c.Param(name)

		if err := v.Var(t, tag); err != nil {
			return err
		}

		return handler(c, t)
	}
}

func HandleBodyWithHeader[T any](handler func(*gin.Context, T) error, v *validator.Validate) func(*gin.Context) error {
	return func(c *gin.Context) error {

		var t T
		if err := c.ShouldBindJSON(&t); err != nil {
			return err
		} else if err = c.ShouldBindHeader(&t); err != nil {
			return err
		} else if err = v.Struct(&t); err != nil {
			return err
		}

		return handler(c, t)
	}
}

func HandleQuery[T any](handler func(*gin.Context, T) error, v *validator.Validate) func(*gin.Context) error {
	return func(c *gin.Context) error {

		var t T
		if err := c.ShouldBindQuery(&t); err != nil {
			return err
		} else if err = v.Struct(&t); err != nil {
			return err
		}

		return handler(c, t)
	}
}
