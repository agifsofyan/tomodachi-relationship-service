package auth

import (
	"context"

	domainError "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/errors"
	"github.com/google/uuid"
)

func UserID(ctx context.Context) (uuid.UUID, error) {

	value := ctx.Value(ContextUserID)
	if value == nil {
		return uuid.Nil, domainError.ErrUnauthorized
	}

	userID, ok := value.(uuid.UUID)
	if !ok {
		return uuid.Nil, domainError.ErrUnauthorized
	}

	return userID, nil
}

func Claims(ctx context.Context) (*JwtClaims, error) {

	value := ctx.Value(ContextClaims)
	if value == nil {
		return nil, domainError.ErrUnauthorized
	}

	claims, ok := value.(*JwtClaims)
	if !ok {
		return nil, domainError.ErrUnauthorized
	}

	return claims, nil
}

func WithUserID(
	ctx context.Context,
	userID uuid.UUID,
) context.Context {

	return context.WithValue(
		ctx,
		ContextUserID,
		userID,
	)
}

func WithClaims(
	ctx context.Context,
	claims *JwtClaims,
) context.Context {

	return context.WithValue(
		ctx,
		ContextClaims,
		claims,
	)
}
