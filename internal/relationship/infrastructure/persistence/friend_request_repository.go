package persistence

import (
	"context"
	"errors"

	domainEntity "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	enum "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/enums"
	domainError "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/errors"
	domainRepository "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/repository"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/mapper"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/model"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type friendRequestRepository struct {
	database.Repository
}

// Create implements [repository.FriendRequestRepository].
func (r *friendRequestRepository) Create(
	ctx context.Context,
	request *domainEntity.FriendRequest,
) error {

	return r.Database(ctx).
		Create(
			mapper.ToFriendRequestModel(request),
		).Error

}

// Delete implements [repository.FriendRequestRepository].
func (r *friendRequestRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	return r.Database(ctx).
		Delete(
			&model.FriendRequestModel{},
			"id = ?",
			id,
		).Error

}

// FindByID implements [repository.FriendRequestRepository].
func (r *friendRequestRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*domainEntity.FriendRequest, error) {

	var m model.FriendRequestModel

	err := r.Database(ctx).
		First(
			&m,
			"id = ?",
			id,
		).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainError.ErrFriendRequestNotFound
		}

		return nil, err
	}

	return mapper.ToFriendRequestEntity(&m), nil

}

// FindPendingBetweenUsers implements [repository.FriendRequestRepository].
func (r *friendRequestRepository) FindPendingBetweenUsers(
	ctx context.Context,
	userA uuid.UUID,
	userB uuid.UUID,
) (*domainEntity.FriendRequest, error) {

	var m model.FriendRequestModel

	err := r.Database(ctx).
		Where(
			`
			(
				requester_id = ?
				AND receiver_id = ?
			)
			OR
			(
				requester_id = ?
				AND receiver_id = ?
			)
			`,
			userA,
			userB,
			userB,
			userA,
		).
		Where(
			"status = ?",
			enum.FriendRequestPending,
		).
		First(&m).
		Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return mapper.ToFriendRequestEntity(&m), nil
}

// FindPendingReceived implements [repository.FriendRequestRepository].
func (r *friendRequestRepository) FindPendingReceived(
	ctx context.Context,
	userID uuid.UUID,
	offset int,
	limit int,
) ([]*domainEntity.FriendRequest, error) {

	var models []model.FriendRequestModel

	err := r.Database(ctx).
		Where(
			"receiver_id = ?",
			userID,
		).
		Where(
			"status = ?",
			enum.FriendRequestPending,
		).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&models).
		Error

	if err != nil {
		return nil, err
	}

	result := make(
		[]*domainEntity.FriendRequest,
		0,
		len(models),
	)

	for _, m := range models {

		modelCopy := m

		result = append(
			result,
			mapper.ToFriendRequestEntity(&modelCopy),
		)

	}

	return result, nil

}

// Update implements [repository.FriendRequestRepository].
func (r *friendRequestRepository) Update(
	ctx context.Context,
	request *domainEntity.FriendRequest,
) error {

	friendRequestModel := mapper.ToFriendRequestModel(request)

	return r.Database(ctx).
		Model(&model.FriendRequestModel{}).
		Where("id = ?", friendRequestModel.ID).
		Updates(friendRequestModel).
		Error
}

func NewFriendRequestRepository(
	db *gorm.DB,
) domainRepository.FriendRequestRepository {

	return &friendRequestRepository{
		Repository: database.NewRepository(db),
	}
}
