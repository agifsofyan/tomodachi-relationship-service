package app

import (
	"fmt"
	"log/slog"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	logger *slog.Logger
	db     *database.Client
	router *gin.Engine
}

func NewServer(
	cfg *config.Config,
	logger *slog.Logger,
	db *database.Client,
	router *gin.Engine,
) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		db:     db,
		router: router,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port)

	s.logger.Info(
		"starting HTTP server",
		slog.String("address", addr),
	)

	return s.router.Run(addr)
}
