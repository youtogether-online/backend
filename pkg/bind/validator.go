package bind

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"reflect"
)

type Valid struct {
	v *validator.Validate
}

func NewValid(v *validator.Validate) *Valid {
	if err := v.RegisterValidation("name", validateName); err != nil {
		logrus.WithError(err).Warn("can't validate name")
	}

	if err := v.RegisterValidation("email", validateEmail); err != nil {
		logrus.WithError(err).Warn("can't validate email")
	}

	if err := v.RegisterValidation("uuid4", validateUUID4); err != nil {
		logrus.WithError(err).Warn("can't validate uuid4")
	}
	return &Valid{v: v}
}

func (v *Valid) Engine() any {
	return v.v
}

func (v *Valid) ValidateStruct(o any) error {
	value := reflect.ValueOf(o)
	valueType := value.Kind()

	if valueType == reflect.Pointer {
		valueType = value.Elem().Kind()
	}

	if valueType == reflect.Struct {
		if err := v.v.Struct(o); err != nil {
			return err
		}
	}
	return nil
}
