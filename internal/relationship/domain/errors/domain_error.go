package errors

import "errors"

var (
	ErrAlreadyFriend = errors.New("already friends")

	ErrRequestAlreadyExists = errors.New("friend request already exists")

	ErrCannotAddSelf = errors.New("cannot add yourself")

	ErrBlocked = errors.New("user is blocked")

	ErrRequestNotFound = errors.New("friend request not found")
)
