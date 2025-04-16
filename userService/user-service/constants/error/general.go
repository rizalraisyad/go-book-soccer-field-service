package error

import "errors"

var (
	ErrInternalServer  = errors.New("Internal Server Error")
	ErrSqlError        = errors.New("Database server failed to execute query")
	ErrTooManyRequests = errors.New("Too many requests")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInvalidToken    = errors.New("invalid token")
	ErrForbidden       = errors.New("forbidden")
)

var GeneralErrors = []error{
	ErrInternalServer,
	ErrSqlError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
