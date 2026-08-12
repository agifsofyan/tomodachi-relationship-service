package app

import (
	"log/slog"
	"net/http"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	cfg *config.Config,
	log *slog.Logger,
	handler *Handlers,
	m *Middlewares,
) *gin.Engine {

	switch cfg.App.Env {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	r.Use(
		middleware.RequestID(),
		middleware.Logger(log),
		middleware.Recovery(log),
		middleware.CORS(),
	)

	_ = r.SetTrustedProxies(nil)

	r.GET("/docs/relationship-service", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/relationship-service/index.html")
	})

	r.GET("/docs/relationship-service/*any", func(c *gin.Context) {
		if c.Param("any") == "/" {
			c.Redirect(http.StatusMovedPermanently, "/docs/relationship-service/index.html")
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

	registerHealth(r, cfg)

	// r.Group("/api/v1")
	registerRelationshipRoutes(r, handler, m)

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

func registerRelationshipRoutes(
	r *gin.Engine,
	handler *Handlers,
	m *Middlewares,
) {

	api := r.Group("/api/v1")

	relationship := api.Group("/relationships")

	relationship.POST(
		"/friends/request",
		m.Auth.Handler(),
		handler.SendFriend.Handle,
	)
}
