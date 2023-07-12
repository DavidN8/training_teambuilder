package handlers

import (
	"net/http"

	"github.com/NahcoCZ/training_teambuilder/backend/models"
	"github.com/NahcoCZ/training_teambuilder/backend/storage"
	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
	DatabaseInstance *storage.Database
}

func (th *TeamHandler) CreateTeam(c echo.Context) error {
	// Create and bind Team from the POST request
	team := &models.Team{}
	if err := c.Bind(team); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Insert Team to database
	// if err := th.DatabaseInstance.InsertTeam(team); err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	c.JSON(http.StatusOK, echo.Map{
		"message":   "Team Created",
		"team_data": team,
	})
	return nil
}

func (th *TeamHandler) DeleteTeam(c echo.Context) error {
	// Handles DELETE request
	// id := c.Param("teamid")

	// Delete Team from database by its ID
	// if err := th.DatabaseInstance.DeleteTeam(id); err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	return nil
}

func (th *TeamHandler) GetTeam(c echo.Context) error {
	id := c.Param("teamid")
	team := &models.Team{}

	// Get id from database
	// if err := th.DatabaseInstance.GetTeam(id, team); err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	c.JSON(http.StatusOK, echo.Map{
		"message":   "GET successful",
		"id":        id,
		"team_data": team,
	})
	return nil
}
