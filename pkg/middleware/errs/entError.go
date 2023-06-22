package errs

import (
	"fmt"
	"net/http"
)

var (
	EntNotFoundError   = newEntError(http.StatusBadRequest, notFound, "Entity is not found")
	EntConstraintError = newEntError(http.StatusBadRequest, alreadyExist, "This value already exists")
)

// entError describes all server-known errors
type entError struct {
	status      int
	code        ErrCode
	Description string
	err         error
}

func newEntError(status int, code ErrCode, description string) entError {
	return entError{status: status, code: code, Description: description}
}

func (e entError) AddError(err error) entError {
	e.err = err
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
		Err:         e.err,
	}
}
