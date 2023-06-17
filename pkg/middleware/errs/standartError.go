package errs

import "net/http"

// StandardError describes all server-known errors
type StandardError struct {
	status      int
	code        ErrCode
	Description string
	err         error
}

// Sign-in errors
var (
	MailCodeError    = newStandardError(http.StatusBadRequest, codeInvalidOrExpired, "Code is not correct, used or expired")
	InvalidPassword  = newStandardError(http.StatusBadRequest, invalidPassword, "Wrong email or password")
	PasswordNotFound = newStandardError(http.StatusBadRequest, passwordNotSet, "You have not registered a password for your account")
)

var (
	UnAuthorized = newStandardError(http.StatusUnauthorized, "", "You are not logged in")
)

// Server errors
var (
	ServerError = newStandardError(http.StatusInternalServerError, serverError, "Server exception was occurred")
	EmailError  = newStandardError(http.StatusInternalServerError, cantSendMail, "Can't send message to your email")
)

// Error implements the Error type
func (e StandardError) Error() string {
	return e.Description
}

// newStandardError creates a new StandardError and returns it
func newStandardError(status int, code ErrCode, description string) StandardError {
	return StandardError{
		status:      status,
		Description: description,
		code:        code,
	}
}

// AddErr saves an error into StandardError and returns it
func (e StandardError) AddErr(err error) StandardError {
	e.err = err
	return e
}

func (e StandardError) GetInfo() *AbstractError {
	return &AbstractError{
		Status:      e.status,
		Code:        e.code,
		Description: e.Description,
		Err:         e.err,
	}
}
