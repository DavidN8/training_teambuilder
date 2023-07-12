package storage

import (
	"github.com/NahcoCZ/training_teambuilder/backend/models"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{
		DB: db,
	}
}

func (db *Database) Migrate() error {
	if err := db.DB.AutoMigrate(&models.User{}, &models.Team{}, &models.Pokemon{}); err != nil {
		return err
	}
	return nil
}

