package mapper

import (
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/model"
)

func ToFriendshipModel(e *entity.Friendship) *model.FriendshipModel {
	if e == nil {
		return nil
	}

	return &model.FriendshipModel{
		ID:        e.ID,
		UserID:    e.UserID,
		FriendID:  e.FriendID,
		CreatedAt: e.CreatedAt,
	}
}

func ToFriendshipEntity(m *model.FriendshipModel) *entity.Friendship {
	if m == nil {
		return nil
	}

	return &entity.Friendship{
		ID:        m.ID,
		UserID:    m.UserID,
		FriendID:  m.FriendID,
		CreatedAt: m.CreatedAt,
	}
}
