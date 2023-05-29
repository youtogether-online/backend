package bind

import (
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/pkg/log"
)

type Valid struct {
	v *validator.Validate
}

func NewValid(v *validator.Validate) *Valid {
	if err := v.RegisterValidation("name", validateName); err != nil {
		log.WithErr(err).Warn("can't validate name")
	}

	if err := v.RegisterValidation("email", validateEmail); err != nil {
		log.WithErr(err).Warn("can't validate email")
	}

	if err := v.RegisterValidation("uuid4", validateUUID4); err != nil {
		log.WithErr(err).Warn("can't validate uuid4")
	}

	if err := v.RegisterValidation("enum", validateEnum); err != nil {
		log.WithErr(err).Warn("can't validate enums")
	}
	return &Valid{v: v}
}

func (v *Valid) Engine() any {
	return v.v
}

func (v *Valid) ValidateStruct(o any) error {
	return v.v.Struct(o)
}
