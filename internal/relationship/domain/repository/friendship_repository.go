package repository

import (
	"context"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	"github.com/google/uuid"
)

type FriendshipRepository interface {
	Create(
		ctx context.Context,
		friendship *entity.Friendship,
	) error

	CreateBatch(
		ctx context.Context,
		friendships []*entity.Friendship,
	) error

	Delete(
		ctx context.Context,
		userID uuid.UUID,
		friendID uuid.UUID,
	) error

	DeleteBothSide(
		ctx context.Context,
		userID uuid.UUID,
		friendID uuid.UUID,
	) error

	Exists(
		ctx context.Context,
		userID uuid.UUID,
		friendID uuid.UUID,
	) (bool, error)

	FindFriends(
		ctx context.Context,
		userID uuid.UUID,
		offset int,
		limit int,
	) ([]*entity.Friendship, error)

	CountFriends(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)

	FindFriendIDs(
		ctx context.Context,
		userID uuid.UUID,
	) ([]uuid.UUID, error)
}
