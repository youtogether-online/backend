package bind

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

var (
	NameRegexp     = regexp.MustCompile(`^[a-zA-Z]\w{3,18}([a-zA-Z0-9])$`)
	EmailRegexp    = regexp.MustCompile(`^\S+@\S+\.\S+$`)
	PasswordRegexp = regexp.MustCompile(`^\w{4,20}$`)
	UUID4          = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`)
)

func validateName(fl validator.FieldLevel) bool {
	return NameRegexp.MatchString(fl.Field().String())
}

func validateEmail(fl validator.FieldLevel) bool {
	return EmailRegexp.MatchString(fl.Field().String())
}

func validateUUID4(fl validator.FieldLevel) bool {
	return UUID4.MatchString(fl.Field().String())
}

func validateEnum(fl validator.FieldLevel) bool {
	enums := strings.Split(fl.Param(), "*")
	field := fl.Field().String()

	for _, v := range enums {
		if v == field {
			return true
		}
	}

	return false
}

func validatePassword(fl validator.FieldLevel) bool {
	return PasswordRegexp.MatchString(fl.Field().String())
}
