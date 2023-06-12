package errs

import "net/http"

// StandardError describes all server-known errors
type StandardError struct {
	status      int
	code        ErrCode
	Description string
	Advice      string
	err         error
}

// Sign-in errors
var (
	MailCodeError    = newStandardError(http.StatusBadRequest, authErr, "Code is not correct", "Try to request a new one")
	PasswordError    = newStandardError(http.StatusBadRequest, authErr, "Wrong email or password", "You can still sign in by your email!")
	PasswordNotFound = newStandardError(http.StatusBadRequest, authErr, "You have not registered a password for your account", "Set it in your profile after authorization")
)

var (
	UnAuthorized = newStandardError(http.StatusUnauthorized, authErr, "You are not logged in", "Click on the button below to sign in!")
)

// Server errors
var (
	ServerError = newStandardError(http.StatusInternalServerError, serverErr, "Server exception was occurred", "Try to restart the page")
	EmailError  = newStandardError(http.StatusInternalServerError, serverErr, "Can't send message to your email", "Try to send it later")
)

// Error implements the Error type
func (e StandardError) Error() string {
	return e.Description
}

// newStandardError creates a new StandardError and returns it
func newStandardError(status int, code ErrCode, description, advice string) StandardError {
	return StandardError{
		status:      status,
		Description: description,
		Advice:      advice,
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
		Advice:      e.Advice,
		Err:         e.err,
	}
}
