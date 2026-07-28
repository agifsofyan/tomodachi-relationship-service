package app

import (
	"log/slog"
	"net/http"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, log *slog.Logger) *gin.Engine {
	// Gin mode
	switch cfg.App.Env {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// Middleware
	r.Use(
		middleware.RequestID(),
		middleware.Logger(log),
		middleware.Recovery(log),
		middleware.CORS(),
	)

	// Trust proxy
	_ = r.SetTrustedProxies(nil)

	// Routes
	registerHealth(r, cfg)

	return r
}

func registerHealth(r *gin.Engine, cfg *config.Config) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "UP",
			"service": cfg.App.Name,
			"version": cfg.App.Version,
		})
	})
}
