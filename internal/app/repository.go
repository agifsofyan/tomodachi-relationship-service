package app

import (
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/repository"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/infrastructure/persistence"
)

type Repositories struct {
	FriendRequest repository.FriendRequestRepository
	Friendship    repository.FriendshipRepository
	Block         repository.BlockRepository
}

func NewRepositories(
	deps *Dependencies,
) *Repositories {

	return &Repositories{

		FriendRequest: persistence.NewFriendRequestRepository(
			deps.Database.DB,
		),

		Friendship: persistence.NewFriendshipRepository(
			deps.Database.DB,
		),

		Block: persistence.NewBlockRepository(
			deps.Database.DB,
		),
	}
}
