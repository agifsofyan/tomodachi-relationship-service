package auth

type ContextKey string

const (
	ContextUserID ContextKey = "user_id"
	ContextClaims ContextKey = "claims"
)
