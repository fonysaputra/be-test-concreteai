package db

import (
	"fmt"
	"log"
	"ms-go-gin/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	var err error

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DsnPostgres,
		PreferSimpleProtocol: true, // Disables prepared statements
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Success connected to database")

	return db
}
