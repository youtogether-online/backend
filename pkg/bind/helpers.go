package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"regexp"
)

var NameRegexp = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{3,18}([a-zA-Z0-9])$")

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
