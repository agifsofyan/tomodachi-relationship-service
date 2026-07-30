package entity

import (
	"time"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/errors"
	"github.com/google/uuid"
)

type Block struct {
	ID uuid.UUID

	BlockerID uuid.UUID

	BlockedID uuid.UUID

	CreatedAt time.Time
}

func NewBlock(
	blockerID uuid.UUID,
	blockedID uuid.UUID,
) (*Block, error) {

	if blockerID == blockedID {
		return nil, errors.ErrCannotBlockSelf
	}

	return &Block{

		ID: uuid.New(),

		BlockerID: blockerID,

		BlockedID: blockedID,

		CreatedAt: time.Now(),
	}, nil

}
