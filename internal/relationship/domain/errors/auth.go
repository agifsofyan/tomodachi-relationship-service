package errors

import "errors"

var (
	ErrUnauthorized            = errors.New("unauthorized")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
)
