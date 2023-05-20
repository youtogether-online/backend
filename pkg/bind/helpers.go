package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"regexp"
)

var (
	NameRegexp  = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,18}([a-zA-Z0-9])$`)
	EmailRegexp = regexp.MustCompile(`^\S+@\S+\.\S+$`)
	UUID4       = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`)
)

// FillStruct of given generic type by request JSON body
func FillStruct[T any](c *gin.Context) (t T, ok bool) {
	if err := c.ShouldBindJSON(&t); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}
	return t, true
}

func validateName(fl validator.FieldLevel) bool {
	return NameRegexp.MatchString(fl.Field().String())
}

func validateEmail(fl validator.FieldLevel) bool {
	return EmailRegexp.MatchString(fl.Field().String())
}

func validateUUID4(fl validator.FieldLevel) bool {
	return UUID4.MatchString(fl.Field().String())
}
