package persistence

import (
	"context"

	"gorm.io/gorm"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	domainRepository "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/repository"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/mapper"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/model"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/database"
	"github.com/google/uuid"
)

type blockRepository struct {
	database.Repository
}

// Create implements [repository.BlockRepository].
func (r *blockRepository) Create(
	ctx context.Context,
	block *entity.Block,
) error {

	return r.Database(ctx).
		Create(
			mapper.ToBlockModel(block),
		).Error

}

// Delete implements [repository.BlockRepository].
func (r *blockRepository) Delete(
	ctx context.Context,
	blockerID uuid.UUID,
	blockedID uuid.UUID,
) error {

	return r.Database(ctx).
		Where("blocker_id = ?", blockerID).
		Where("blocked_id = ?", blockedID).
		Delete(&model.BlockModel{}).
		Error

}

// Exists implements [repository.BlockRepository].
func (r *blockRepository) Exists(
	ctx context.Context,
	blockerID uuid.UUID,
	blockedID uuid.UUID,
) (bool, error) {

	var count int64

	err := r.Database(ctx).
		Model(&model.BlockModel{}).
		Where("blocker_id = ?", blockerID).
		Where("blocked_id = ?", blockedID).
		Count(&count).
		Error

	return count > 0, err

}

// FindBlockedIDs implements [repository.BlockRepository].
func (r *blockRepository) FindBlockedIDs(
	ctx context.Context,
	blockerID uuid.UUID,
) ([]uuid.UUID, error) {

	var ids []uuid.UUID

	err := r.Database(ctx).
		Model(&model.BlockModel{}).
		Where("blocker_id = ?", blockerID).
		Pluck("blocked_id", &ids).
		Error

	if err != nil {
		return nil, err
	}

	return ids, nil

}

// FindBlockedUsers implements [repository.BlockRepository].
func (r *blockRepository) FindBlockedUsers(
	ctx context.Context,
	blockerID uuid.UUID,
	offset int,
	limit int,
) ([]*entity.Block, error) {

	var models []model.BlockModel

	err := r.Database(ctx).
		Where("blocker_id = ?", blockerID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&models).
		Error

	if err != nil {
		return nil, err
	}

	result := make([]*entity.Block, 0, len(models))

	for _, m := range models {

		modelCopy := m

		result = append(
			result,
			mapper.ToBlockEntity(&modelCopy),
		)

	}

	return result, nil

}

// IsBlockedBetweenUsers implements [repository.BlockRepository].
func (r *blockRepository) IsBlockedBetweenUsers(
	ctx context.Context,
	userA uuid.UUID,
	userB uuid.UUID,
) (bool, error) {

	var count int64

	err := r.Database(ctx).
		Model(&model.BlockModel{}).
		Where(
			"(blocker_id = ? AND blocked_id = ?) OR (blocker_id = ? AND blocked_id = ?)",
			userA,
			userB,
			userB,
			userA,
		).
		Count(&count).
		Error

	return count > 0, err

}

func NewBlockRepository(
	db *gorm.DB,
) domainRepository.BlockRepository {

	return &blockRepository{
		Repository: database.NewRepository(db),
	}
}
