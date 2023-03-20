package exceptions

// MyError simplify the error handling stage. All native errors must be this type
type MyError struct {
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"message,omitempty"`
	Advice string `json:"advice,omitempty"`
	Err    error  `json:"err,omitempty"`
}

// Error implements the Error type
func (e MyError) Error() string {
	return e.Err.Error()
}

// newError creates a new MyError and returns it
func newError(code int, msg, advice string) MyError {
	return MyError{
		Msg:    msg,
		Code:   code,
		Advice: advice,
	}
}

// AddErr saves an error into MyError and returns it
func (e MyError) AddErr(err error) MyError {
	e.Err = err
	return e
}
