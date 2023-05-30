package errs

// EntError describes all server-known errors
type EntError struct {
	Status int               `json:"-"`
	Msg    string            `json:"message,omitempty"`
	Fields map[string]string `json:"fields,omitempty"`
	Advice string            `json:"advice,omitempty"`
	Err    error             `json:"-"`
}

// Error implements the Error type
func (e EntError) Error() string {
	return e.Msg
}

func (e EntError) GetInfo() *AbstractError {
	var msg any
	if e.Msg != "" {
		msg = e.Msg
	} else {
		msg = e.Fields
	}

	return &AbstractError{
		Status: e.Status,
		Msg:    msg,
		Advice: e.Advice,
		Err:    e.Err,
	}
}

// newEntError creates a new EntError and returns it
func newEntError() EntError {
	return EntError{} //TODO
}
