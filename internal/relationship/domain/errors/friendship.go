package errors

import "errors"

var (
	ErrAlreadyFriend  = errors.New("users are already friends")
	ErrFriendNotFound = errors.New("friendship not found")
)
