package app

import (
	"log"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/database"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/logger"

	_ "github.com/agifsofyan/tomodachi-relationship-service/docs" // generated docs, wajib di-import
)

// @title           Relationship Service API
// @version         1.0
// @description     API documentation for Relationship Service
// @BasePath        /
func Run() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log := logger.New(cfg)

	db, err := database.New(cfg, log)
	if err != nil {
		log.Error("failed to connect database")
		panic(err)
	}

	deps := NewDependencies(
		cfg,
		db,
	)

	repositories := NewRepositories(deps)

	services := NewServices(
		deps,
		repositories,
	)

	handlers := NewHandlers(
		services,
	)

	middlewares := NewMiddlewares(cfg)

	router := NewRouter(
		cfg,
		log,
		handlers,
		middlewares,
	)

	server := NewServer(
		cfg,
		log,
		db,
		router,
	)

	if err := server.Run(); err != nil {
		log.Error(err.Error())
	}
}
