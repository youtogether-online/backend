package errs

import "net/http"

// Database errors (templates)
var (
	RedisNilError = newRedisError(http.StatusBadRequest, notFoundErr, "Can't find value", "Register, please")
	RedisTxError  = newRedisError(http.StatusInternalServerError, serverErr, "Operation failed", "Try to request it later")
)

// redisError describes all server-known errors
type redisError struct {
	status      int
	code        ErrCode
	Description string
	Advice      string
	err         error
}

func newRedisError(status int, code ErrCode, description string, advice string) redisError {
	return redisError{status: status, code: code, Description: description, Advice: advice}
}

func (r redisError) AddError(err error) redisError {
	r.err = err
	return r
}

// Error implements the Error type
func (r redisError) Error() string {
	return r.Description
}

func (r redisError) GetInfo() *AbstractError {

	return &AbstractError{
		Status:      r.status,
		Code:        r.code,
		Description: r.Description,
		Advice:      r.Advice,
		Err:         r.err,
	}
}
