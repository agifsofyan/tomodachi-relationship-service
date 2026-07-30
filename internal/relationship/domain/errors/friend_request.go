package errors

import "errors"

var (
	ErrFriendRequestNotFound      = errors.New("friend request not found")
	ErrCannotRequestSelf          = errors.New("cannot send friend request to yourself")
	ErrFriendRequestAlreadyExists = errors.New("friend request already exists")
)
