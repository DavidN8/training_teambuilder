package handlers

import (
	"net/http"

	"github.com/NahcoCZ/training_teambuilder/backend/models"
	"github.com/NahcoCZ/training_teambuilder/backend/storage"
	"github.com/labstack/echo/v4"
)

type PokemonHandler struct {
	DatabaseInstance *storage.Database
}

func (ph *PokemonHandler) CreatePokemon(c echo.Context) error {
	pokemon := &models.Pokemon{}

	if err := c.Bind(pokemon); err != nil {
		return err
	}
	return nil
}

func (ph *PokemonHandler) GetPokemon(c echo.Context) error {
	// id := c.Param("pokemonid")
	pokemon := &models.Pokemon{}

	// Get Pokemon details from the Database
	// if err := ph.DatabaseInstance.GetPokemon(id, pokemon); err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	c.JSON(http.StatusOK, echo.Map{
		"data": pokemon,
	})
	return nil
}

func (ph *PokemonHandler) DeletePokemon(c echo.Context) error {
	// id := c.Param("pokemonid")

	return nil
}
