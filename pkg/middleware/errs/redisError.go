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
	description string
}

func newRedisError(status int, code ErrCode, description string) redisError {
	return redisError{status: status, code: code, description: description}
}

// Error implements the Error type
func (r redisError) Error() string {
	return r.description
}

func (r redisError) GetInfo() *AbstractError {

	return &AbstractError{
		Status:      r.status,
		Code:        r.code,
		Description: r.description,
	}
}
