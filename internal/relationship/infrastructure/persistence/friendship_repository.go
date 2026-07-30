package persistence

import (
	"context"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	domainRepository "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/repository"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/mapper"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/model"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type friendshipRepository struct {
	database.Repository
}

// CountFriends implements [repository.FriendshipRepository].
func (r *friendshipRepository) CountFriends(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	var count int64

	err := r.Database(ctx).
		Model(&model.FriendshipModel{}).
		Where("user_id = ?", userID).
		Count(&count).
		Error

	return count, err

}

// Create implements [repository.FriendshipRepository].
func (r *friendshipRepository) Create(
	ctx context.Context,
	friendship *entity.Friendship,
) error {

	return r.Database(ctx).
		Create(
			mapper.ToFriendshipModel(friendship),
		).Error

}

// CreateBatch implements [repository.FriendshipRepository].
func (r *friendshipRepository) CreateBatch(
	ctx context.Context,
	friendships []*entity.Friendship,
) error {

	models := make([]*model.FriendshipModel, 0, len(friendships))

	for _, friendship := range friendships {
		models = append(
			models,
			mapper.ToFriendshipModel(friendship),
		)
	}

	return r.Database(ctx).
		Create(&models).
		Error

}

// Delete implements [repository.FriendshipRepository].
func (r *friendshipRepository) Delete(
	ctx context.Context,
	userID uuid.UUID,
	friendID uuid.UUID,
) error {

	return r.Database(ctx).
		Where("user_id = ?", userID).
		Where("friend_id = ?", friendID).
		Delete(&model.FriendshipModel{}).
		Error

}

// DeleteBothSide implements [repository.FriendshipRepository].
func (r *friendshipRepository) DeleteBothSide(
	ctx context.Context,
	userID uuid.UUID,
	friendID uuid.UUID,
) error {

	return r.Database(ctx).
		Where(
			"(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
			userID,
			friendID,
			friendID,
			userID,
		).
		Delete(&model.FriendshipModel{}).
		Error

}

// Exists implements [repository.FriendshipRepository].
func (r *friendshipRepository) Exists(
	ctx context.Context,
	userID uuid.UUID,
	friendID uuid.UUID,
) (bool, error) {

	var count int64

	err := r.Database(ctx).
		Model(&model.FriendshipModel{}).
		Where("user_id = ?", userID).
		Where("friend_id = ?", friendID).
		Count(&count).
		Error

	return count > 0, err

}

// FindFriendIDs implements [repository.FriendshipRepository].
func (r *friendshipRepository) FindFriendIDs(
	ctx context.Context,
	userID uuid.UUID,
) ([]uuid.UUID, error) {

	var ids []uuid.UUID

	err := r.Database(ctx).
		Model(&model.FriendshipModel{}).
		Where("user_id = ?", userID).
		Pluck("friend_id", &ids).
		Error

	if err != nil {
		return nil, err
	}

	return ids, nil

}

// FindFriends implements [repository.FriendshipRepository].
func (r *friendshipRepository) FindFriends(
	ctx context.Context,
	userID uuid.UUID,
	offset int,
	limit int,
) ([]*entity.Friendship, error) {

	var models []model.FriendshipModel

	err := r.Database(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&models).
		Error

	if err != nil {
		return nil, err
	}

	friendships := make([]*entity.Friendship, 0, len(models))

	for _, m := range models {

		friendship := mapper.ToFriendshipEntity(&m)

		friendships = append(friendships, friendship)

	}

	return friendships, nil

}

func NewFriendshipRepository(
	db *gorm.DB,
) domainRepository.FriendshipRepository {

	return &friendshipRepository{
		Repository: database.NewRepository(db),
	}
}
