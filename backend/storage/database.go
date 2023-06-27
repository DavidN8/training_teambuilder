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

func (db *Database) CreateUser(data *models.User) error {
	if err := db.DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) GetUser(id string, data *models.User) error {
	if err := db.DB.Find(data, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) DeleteUser(id string) error {
	// Find value to delete
	if err := db.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
