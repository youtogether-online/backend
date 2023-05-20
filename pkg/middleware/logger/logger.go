package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
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

type PlainFormatter struct {
	TimestampFormat string
}

func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s |%s\n", entry.Time.Format(
		f.TimestampFormat), entry.Message)), nil
}

type QueryHandler struct {
	log *logrus.Logger
}

func NewQueryHandler(log *logrus.Logger) *QueryHandler {
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&PlainFormatter{TimestampFormat: "2006/01/02 15:32:05"})
	log.SetOutput(os.Stdout)
	return &QueryHandler{log: log}
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
