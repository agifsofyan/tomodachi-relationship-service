package main

import (
	"fmt"
	"log"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("migration database:", cfg.Database.Name)

	fmt.Print(cfg.Database.MigrationURL())
}
