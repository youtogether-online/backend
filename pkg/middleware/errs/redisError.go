package errs

// RedisError describes all server-known errors
type RedisError struct {
	Status int    `json:"-"`
	Msg    string `json:"message,omitempty"`
	Advice string `json:"advice,omitempty"`
	Err    error  `json:"-"`
}

func NewRedisError(status int, msg string, advice string) *RedisError {
	return &RedisError{Status: status, Msg: msg, Advice: advice}
}

func (r RedisError) AddError(err error) RedisError {
	r.Err = err
	return r
}

// Error implements the Error type
func (r RedisError) Error() string {
	return r.Msg
}

func (r RedisError) GetInfo() *AbstractError {

	return &AbstractError{
		Status: r.Status,
		Msg:    r.Msg,
		Advice: r.Advice,
		Err:    r.Err,
	}
}
