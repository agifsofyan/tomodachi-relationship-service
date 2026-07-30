package errors

import "errors"

var (
	ErrCannotBlockSelf = errors.New("cannot block yourself")
	ErrAlreadyBlocked  = errors.New("user already blocked")
	ErrUserBlocked     = errors.New("user is blocked")
)
