package errs

import "net/http"

// Database errors (templates)
var (
	RedisNilError = newRedisError(http.StatusBadRequest, notFound, "Can't find value")
	RedisTxError  = newRedisError(http.StatusInternalServerError, txFailed, "Operation failed")
	RedisError    = newRedisError(http.StatusInternalServerError, serverError, "Can't perform query")
)

// redisError describes all server-known errors
type redisError struct {
	status      int
	code        ErrCode
	Description string
	err         error
}

func newRedisError(status int, code ErrCode, description string) redisError {
	return redisError{status: status, code: code, Description: description}
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
		Err:         r.err,
	}
}
