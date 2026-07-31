package mapper

import (
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/model"
)

func ToBlockModel(e *entity.Block) *model.BlockModel {
	if e == nil {
		return nil
	}

	return &model.BlockModel{
		ID:        e.ID,
		BlockerID: e.BlockerID,
		BlockedID: e.BlockedID,
		CreatedAt: e.CreatedAt,
	}
}

func ToBlockEntity(m *model.BlockModel) *entity.Block {
	if m == nil {
		return nil
	}

	return &entity.Block{
		ID:        m.ID,
		BlockerID: m.BlockerID,
		BlockedID: m.BlockedID,
		CreatedAt: m.CreatedAt,
	}
}
