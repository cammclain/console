package database

import (
	"fmt"
	"log"
	"server/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres
func ConnectToPostgres(config *config.Settings) (*gorm.DB, error) {
	// Build the Postgres connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDBName,
		config.PostgresPort,
	)

	// Connect to the Postgres database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Default().Println("Error connecting to Postgres:", err)
		return nil, err
	}
	// Return the Postgres client connection
	return db, nil
}
