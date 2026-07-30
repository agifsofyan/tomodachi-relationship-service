package enum

type FriendRequestStatus string

const (
	FriendRequestPending   FriendRequestStatus = "PENDING"
	FriendRequestAccepted  FriendRequestStatus = "ACCEPTED"
	FriendRequestRejected  FriendRequestStatus = "REJECTED"
	FriendRequestCancelled FriendRequestStatus = "CANCELLED"
)

func (s FriendRequestStatus) IsValid() bool {
	switch s {
	case FriendRequestPending,
		FriendRequestAccepted,
		FriendRequestRejected,
		FriendRequestCancelled:
		return true
	default:
		return false
	}
}
