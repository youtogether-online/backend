package handleQuery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"time"
)

const (
	red    string = "\033[41m"
	green  string = "\033[42m"
	yellow string = "\033[43m"
	blue   string = "\033[44m"
	cyan   string = "\033[46m"
	white  string = "\033[47m"
)

type QueryHandler struct {
	log *log.Logger
}

func NewQueryHandler() *QueryHandler {
	return &QueryHandler{log: log.NewLogger(log.InfoLevel, &log.TextFormatter{}, false)}
}

func (q *QueryHandler) HandleQueries(c *gin.Context) {
	s := time.Now()
	c.Next()
	l := time.Since(s)

	defaultStatus, status := c.Writer.Status(), ""
	switch defaultStatus / 100 {
	case 2:
		status = setColor(defaultStatus, green)
	case 3:
		status = setColor(defaultStatus, blue)
	case 4:
		status = setColor(defaultStatus, yellow)
	case 5:
		status = setColor(defaultStatus, red)
	default:
		status = setColor(defaultStatus, white)
	}

	method := c.Request.Method
	switch method {
	case "POST":
		method = setColor(method, green)
	case "GET":
		method = setColor(method, blue)
	case "PATCH":
		method = setColor(method, cyan)
	case "DELETE":
		method = setColor(method, red)
	default:
		method = setColor(method, white)
	}

	q.log.Infof("%s| %12v |%6s| %s\n", status, l, method, c.Request.RequestURI)
}

func setColor(text any, color string) string {
	return fmt.Sprintf("%s %v \033[0m", color, text)
}
