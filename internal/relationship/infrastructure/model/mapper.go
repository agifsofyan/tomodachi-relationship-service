package model

import (
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
)

func ToFriendRequestEntity(m *FriendRequestModel) *entity.FriendRequest {
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

func ToFriendRequestModel(e *entity.FriendRequest) *FriendRequestModel {
	if e == nil {
		return nil
	}

	return &FriendRequestModel{
		ID:          e.ID,
		RequesterID: e.RequesterID,
		ReceiverID:  e.ReceiverID,
		Status:      e.Status,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
