package errs

// StandardError describes all server-known errors
type StandardError struct {
	Status int    `json:"-"`
	Msg    string `json:"message,omitempty"`
	Advice string `json:"advice,omitempty"`
	Err    error  `json:"-"`
}

// Error implements the Error type
func (e StandardError) Error() string {
	return e.Msg
}

func (e StandardError) GetInfo() *AbstractError {
	return &AbstractError{
		Status: e.Status,
		Msg:    e.Msg,
		Advice: e.Advice,
		Err:    e.Err,
	}
}

// newStandardError creates a new StandardError and returns it
func newStandardError(status int, msg, advice string) StandardError {
	return StandardError{
		Status: status,
		Msg:    msg,
		Advice: advice,
	}
}

// AddErr saves an error into StandardError and returns it
func (e StandardError) AddErr(err error) StandardError {
	e.Err = err
	return e
}
