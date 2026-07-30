package repository

import (
	"context"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	"github.com/google/uuid"
)

type BlockRepository interface {
	Create(
		ctx context.Context,
		block *entity.Block,
	) error

	Delete(
		ctx context.Context,
		blockerID uuid.UUID,
		blockedID uuid.UUID,
	) error

	Exists(
		ctx context.Context,
		blockerID uuid.UUID,
		blockedID uuid.UUID,
	) (bool, error)

	FindBlockedUsers(
		ctx context.Context,
		blockerID uuid.UUID,
		offset int,
		limit int,
	) ([]*entity.Block, error)

	FindBlockedIDs(
		ctx context.Context,
		blockerID uuid.UUID,
	) ([]uuid.UUID, error)

	IsBlockedBetweenUsers(
		ctx context.Context,
		userA uuid.UUID,
		userB uuid.UUID,
	) (bool, error)
}
