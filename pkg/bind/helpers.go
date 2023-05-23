package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"regexp"
	"strings"
)

var (
	NameRegexp  = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,18}([a-zA-Z0-9])$`)
	EmailRegexp = regexp.MustCompile(`^\S+@\S+\.\S+$`)
	UUID4       = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`)
)

/*
// FillStructJSON of given generic type by request JSON body
func FillStructJSON[T any](c *gin.Context) (t T) {
	h := c.Request.Header
	b := http.MaxBytesReader(c.Writer, c.Request.Body, 1048576) // 1 MB

	if b == nil || h.Get("Content-Type") != "application/json" {
		logrus.Warn("invalid Content-Type")
		return t
	}

	dec := json.NewDecoder(b)
	dec.DisallowUnknownFields()
	err := dec.Decode(&t)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {

		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			status := http.StatusBadRequest

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			status := http.StatusBadRequest

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			status := http.StatusBadRequest
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			status := http.StatusBadRequest

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			status := http.StatusBadRequest

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			status := http.StatusRequestEntityTooLarge

		default:
			logrus.WithError(err).Warn("srv err")
			return nil
		}
		return
	}

	err = dec.Decode(&struct{}{})
	if err = dec.Decode(&struct{}{}); err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		status := http.StatusBadRequest
		return
	}
	return
}

*/

// FillStructJSON of given generic type by request JSON body and headers. Header has authority on body
func FillStructJSON[T any](c *gin.Context) (t *T) {
	if err := c.ShouldBindJSON(t); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}
	return
}

// FillStructByHeader of given generic type by request JSON body and headers. Header has authority on body
func FillStructByHeader[T any](c *gin.Context) (t T) {
	if err := c.ShouldBindHeader(&t); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}
	return
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
