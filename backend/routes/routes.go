package routes

import (
	"github.com/NahcoCZ/training_teambuilder/backend/handlers"
	"github.com/NahcoCZ/training_teambuilder/backend/storage"
	"github.com/labstack/echo/v4"
)

func CreateRoutes(e *echo.Echo, db *storage.Database) {
	userHandler := &handlers.UserHandler{DatabaseInstance: db}
	teamHandler := &handlers.TeamHandler{DatabaseInstance: db}
	pokemonHandler := &handlers.PokemonHandler{DatabaseInstance: db}
	setupUserRoutes(e, userHandler)
	setupTeamRoutes(e, teamHandler)
	setupPokemonRoutes(e, pokemonHandler)
}

func setupUserRoutes(e *echo.Echo, handler *handlers.UserHandler) {
	e.POST("/user", handler.CreateUser)
	e.DELETE("/user/:userid", handler.DeleteUser)
	e.GET("/user/:userid", handler.GetUser)
	e.GET("/user", handler.GetAllUsers)
}

func setupTeamRoutes(e *echo.Echo, handler *handlers.TeamHandler) {
	e.POST("/team", handler.CreateTeam)
	e.DELETE("/team/:teamid", handler.DeleteTeam)
	e.GET("/team/:teamid", handler.GetTeam)
}

func setupPokemonRoutes(e *echo.Echo, handler *handlers.PokemonHandler) {
	e.POST("/pokemon", handler.CreatePokemon)
	e.DELETE("/pokemon/:pokemonid", handler.DeletePokemon)
	e.GET("/pokemon/:pokemonid", handler.GetPokemon)
}
