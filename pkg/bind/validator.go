package bind

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type Valid struct {
	v *validator.Validate
}

func NewValid(v *validator.Validate) *Valid {
	v.RegisterValidation("name", validateName)
	return &Valid{v: v}
}

func validateName(fl validator.FieldLevel) bool {
	return NameRegexp.MatchString(fl.Field().String())
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
