package mapper

import (
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/model"
)

func ToFriendRequestModel(e *entity.FriendRequest) *model.FriendRequestModel {
	if e == nil {
		return nil
	}

	return &model.FriendRequestModel{
		ID:          e.ID,
		RequesterID: e.RequesterID,
		ReceiverID:  e.ReceiverID,
		Status:      e.Status,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToFriendRequestEntity(m *model.FriendRequestModel) *entity.FriendRequest {
	if m == nil {
		return nil
	}

	return &entity.FriendRequest{
		ID:          m.ID,
		RequesterID: m.RequesterID,
		ReceiverID:  m.ReceiverID,
		Status:      m.Status,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
