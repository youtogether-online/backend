package errs

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// ValidError describes all validation errors
type ValidError struct {
	Status int               `json:"-"`
	Fields map[string]string `json:"fields,omitempty"`
}

func (e *ValidError) GetInfo() *AbstractError {
	return &AbstractError{
		Status: e.Status,
		Msg:    e.Fields,
		Advice: "",
		Err:    nil,
	}
}

// newValidError creates a new StandardError and returns it
func newValidError(errs validator.ValidationErrors) *ValidError {
	fields := make(map[string]string)

	for _, err := range errs {
		field := err.Field()
		switch err.Tag() {
		case "email":
			fields[field] = fmt.Sprintf("%s is not the correct email", field)
		case "required":
			fields[field] = fmt.Sprintf("%s should not be empty", field)
		case "numeric":
			fields[field] = fmt.Sprintf("%s must be a number", field)
		case "gte":
			fields[field] = fmt.Sprintf("%s must be greater or equal %s", field, err.Param())
		case "lte":
			fields[field] = fmt.Sprintf("%s must be lesser or equal %s", field, err.Param())
		case "len":
			fields[field] = fmt.Sprintf("%s must have a length of %s", field, err.Param())
		case "gt":
			fields[field] = fmt.Sprintf("%s must be greater than %s", field, err.Param())
		case "lt":
			fields[field] = fmt.Sprintf("%s must be lesser than %s", field, err.Param())
		case "printascii":
			fields[field] = fmt.Sprintf("%s can contain only /:@-._~!?$&'()*+,;= and any english letters", field)
		case "name":
			fields[field] = fmt.Sprintf("%s is not valid name", field)
		case "jwt":
			fields[field] = fmt.Sprintf("%s is not jwt format", field)
		case "title":
			fields[field] = fmt.Sprintf("%s is not a title", field)
		case "inn":
			fields[field] = fmt.Sprintf("%s is not an INN", field)
		case "link":
			fields[field] = fmt.Sprintf("%s is not a link", field)
		case "password":
			fields[field] = fmt.Sprintf("%s must contain only english letters and _ character", field)
		case "enum":
			fields[field] = fmt.Sprintf("%s can have %s values", field, strings.ReplaceAll(err.Param(), "*", ", "))
		}
	}

	return &ValidError{
		Status: http.StatusBadRequest,
		Fields: fields,
	}
}