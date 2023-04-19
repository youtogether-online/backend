package exceptions

// @Description All native errors must be this type
type MyError struct {
	Status int    `json:"-"`
	Msg    string `json:"message,omitempty" example:"Exception was occurred"`
	Advice string `json:"advice,omitempty" example:"Try to send request later"`
	Err    error  `json:"-"`
}

// Error implements the Error type
func (e MyError) Error() string {
	return e.Msg
}

// newError creates a new MyError and returns it
func newError(status int, msg, advice string) MyError {
	return MyError{
		Status: status,
		Msg:    msg,
		Advice: advice,
	}
}

// AddErr saves an error into MyError and returns it
func (e MyError) AddErr(err error) MyError {
	e.Err = err
	return e
}
