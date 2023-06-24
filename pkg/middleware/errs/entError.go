package errs

import (
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
	description string
}

func newEntError(status int, code ErrCode, description string) *entError {
	return &entError{status: status, code: code, description: description}
}

// Error implements the Error type
func (e entError) Error() string {
	return e.description
}

func (e entError) GetInfo() *AbstractError {
	return &AbstractError{
		Status:      e.status,
		Code:        e.code,
		Description: e.description,
	}
}
