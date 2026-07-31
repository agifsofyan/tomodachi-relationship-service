package app

import (
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/auth"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/middleware"
)

type Middlewares struct {
	Auth *middleware.AuthMiddleware
}

func NewMiddlewares(
	cfg *config.Config,
) *Middlewares {

	parser := auth.NewParser(
		cfg.JWT.SecretKey,
		cfg.JWT.Algorithm,
	)

	return &Middlewares{

		Auth: middleware.NewAuthMiddleware(
			parser,
		),
	}
}
