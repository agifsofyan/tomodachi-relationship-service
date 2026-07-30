package sendfriendrequest

import (
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/repository"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/database"
)

type SendFriendServiceDependencies struct {
	Transaction             database.Transaction
	FriendRequestRepository repository.FriendRequestRepository
	FriendshipRepository    repository.FriendshipRepository
	BlockRepository         repository.BlockRepository
}
