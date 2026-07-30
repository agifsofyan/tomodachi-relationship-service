package repository

import (
	"context"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	"github.com/google/uuid"
)

type FriendRequestRepository interface {
	Create(ctx context.Context, request *entity.FriendRequest) error
	Update(ctx context.Context, request *entity.FriendRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.FriendRequest, error)
	FindPendingBetweenUsers(ctx context.Context, requesterID uuid.UUID, receiverID uuid.UUID) (*entity.FriendRequest, error)
	FindPendingReceived(ctx context.Context, userID uuid.UUID, offset int, limit int) ([]*entity.FriendRequest, error)
}
