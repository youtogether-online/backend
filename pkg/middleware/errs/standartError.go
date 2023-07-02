package errs

import "net/http"

// StandardError describes all server-known errors
type StandardError struct {
	status      int
	code        ErrCode
	description string
	err         error
}

// Sign-in errors
var (
	MailCodeError    = newStandardError(http.StatusBadRequest, codeInvalidOrExpired, "Code is not correct, used or expired")
	InvalidPassword  = newStandardError(http.StatusBadRequest, invalidPassword, "Wrong email or password")
	PasswordNotFound = newStandardError(http.StatusBadRequest, passwordNotSet, "You have not registered a password for your account")
)

// Web Socket errors
var (
	WebsocketNotFound = newStandardError(http.StatusBadRequest, websocketExcepted, "Web socket connection excepted")
)

// File errors
var (
	UnsupportedImageType = newStandardError(http.StatusUnsupportedMediaType, "", "Image should have .png, .jpg or .jpeg extension")
	ImageTooLarge        = newStandardError(http.StatusRequestEntityTooLarge, "", "Image must be smaller than 3000 by 3000 pixels")
)

// UnAuthorized errors
var UnAuthorized = newStandardError(http.StatusUnauthorized, "", "You are not logged in")

// Server errors
var (
	EmailError = newStandardError(http.StatusInternalServerError, cantSendMail, "Can't send message to your email")
)

// Error implements the Error type
func (e StandardError) Error() string {
	return e.err.Error()
}

// newStandardError creates a new StandardError and returns it
func newStandardError(status int, code ErrCode, description string) StandardError {
	return StandardError{
		status:      status,
		description: description,
		code:        code,
	}
}

func (e StandardError) AddErr(err error) StandardError {
	e.err = err
	return e
}

func (e StandardError) GetInfo() *AbstractError {
	return &AbstractError{
		Status:      e.status,
		Code:        e.code,
		Description: e.description,
		Err:         e.err,
	}
}
