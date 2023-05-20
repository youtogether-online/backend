package errs

// MyError describes all server-known errors
type MyError struct {
	Status int    `json:"-"`
	Msg    string `json:"message,omitempty"`
	Advice string `json:"advice,omitempty"`
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
