package enums

type FriendRequestStatus string

const (
	RequestPending  FriendRequestStatus = "PENDING"
	RequestAccepted FriendRequestStatus = "ACCEPTED"
	RequestRejected FriendRequestStatus = "REJECTED"
	RequestCanceled FriendRequestStatus = "CANCELED"
)
