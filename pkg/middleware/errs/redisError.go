package errs

// RedisError describes all server-known errors
type RedisError struct {
	Status int               `json:"-"`
	Msg    string            `json:"message,omitempty"`
	Fields map[string]string `json:"fields,omitempty"`
	Advice string            `json:"advice,omitempty"`
	Err    error             `json:"-"`
}

// Error implements the Error type
func (e RedisError) Error() string {
	return e.Msg
}

func (e RedisError) GetInfo() *AbstractError {
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

// newRedisError creates a new RedisError and returns it
func newRedisError() RedisError {
	return RedisError{} //TODO
}
