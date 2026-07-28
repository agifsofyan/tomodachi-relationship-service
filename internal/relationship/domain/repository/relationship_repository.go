package repository

import "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"

type RelationshipRepository interface {
	CreateFriendRequest(
		request *entity.FriendRequest,
	) error

	GetFriendRequest(
		requestID string,
	) (*entity.FriendRequest, error)

	UpdateFriendRequest(
		request *entity.FriendRequest,
	) error

	IsFriend(
		userID string,
		friendID string,
	) (bool, error)

	IsBlocked(
		userID string,
		friendID string,
	) (bool, error)

	CreateFriendship(
		friendship *entity.Friendship,
	) error
}
