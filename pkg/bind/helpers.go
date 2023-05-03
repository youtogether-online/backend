package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/internal/middleware/exceptions"
	"regexp"
)

var NameRegexp = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{3,18}([a-zA-Z0-9])$")

func FillStruct[T dto.FilledDTO](c *gin.Context) (t T, ok bool) {
	if err := c.ShouldBindJSON(&t); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}
	return t, true
}
