package errs

import (
	"fmt"
	"net/http"
)

var (
	EntNotFoundError    = newEntError(http.StatusBadRequest, notFoundErr, "Not found", "Try to find another entity")
	EntValidError       = newEntError(http.StatusBadRequest, validErr, "", "")
	EntConstraintError  = newEntError(http.StatusBadRequest, notFoundErr, "Can't set this value", "Try to set constraint value")
	EntNotSingularError = newEntError(http.StatusInternalServerError, serverErr, "Single result was expected, but several were found", "Try to look for something else")
	EntNotLoadedError   = newEntError(http.StatusInternalServerError, serverErr, "Can't load data", "Try to request it later")
)

// entError describes all server-known errors
type entError struct {
	status      int
	code        ErrCode
	Description string
	Fields      map[string]string
	Advice      string
	err         error
}

func newEntError(status int, code ErrCode, description, advice string) entError {
	return entError{status: status, code: code, Description: description, Advice: advice}
}

func (e entError) AddError(err error) entError {
	e.err = err
	return e
}

func (e entError) SetFields(fields map[string]string) entError {
	e.Fields = fields
	return e
}

func (e entError) SetDB(dbName string) entError {
	e.Description = fmt.Sprintf(e.Description, dbName)
	return e
}

// Error implements the Error type
func (e entError) Error() string {
	return e.Description
}

func (e entError) GetInfo() *AbstractError {
	return &AbstractError{
		Status:      e.status,
		Code:        e.code,
		Description: e.Description,
		Fields:      e.Fields,
		Advice:      e.Advice,
		Err:         e.err,
	}
}
