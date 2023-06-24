package storage

import (
	"fmt"

	"github.com/NahcoCZ/training_teambuilder/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	// Create Data Source Name from config passed in the parameter
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode,
	)

	// Open a gorm.DB through the postgres server with the given dsn
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
