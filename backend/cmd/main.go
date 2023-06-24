package main

import (
	"log"

	"github.com/NahcoCZ/training_teambuilder/backend/config"
	"github.com/NahcoCZ/training_teambuilder/backend/routes"
	"github.com/NahcoCZ/training_teambuilder/backend/storage"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load Environment Variables from .env
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Can't load environment variables:", err)
	}

	// Get DB Configuration
	db_config := config.LoadConfig()

	// Create postgres connection using the config
	db, err := storage.ConnectDatabase(db_config)
	if err != nil {
		log.Fatal("Can't connect database:", err)
	}
	databaseInstance := storage.NewDatabase(db)
	if err := databaseInstance.Migrate(); err != nil {
		log.Fatal("Could not migrate database:", err)
	}

	// Create Backend Server and establish routes
	e := echo.New()
	routes.CreateRoutes(e, databaseInstance)
}
