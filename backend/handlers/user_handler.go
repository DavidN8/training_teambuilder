package handlers

import (
	"net/http"

	"github.com/NahcoCZ/training_teambuilder/backend/models"
	"github.com/NahcoCZ/training_teambuilder/backend/storage"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	DatabaseInstance *storage.Database
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	// Binding User from POST request
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Inputting into database
	if err := uh.DatabaseInstance.CreateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, echo.Map{
		"message": "User Creation Successful",
	})
	return nil
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("userid")
	user := &models.User{}

	// Get User data from Database
	if err := uh.DatabaseInstance.GetUser(id, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
	return nil
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("userid")

	// Delete User data from Database
	if err := uh.DatabaseInstance.DeleteUser(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, echo.Map{
		"message": "User Deletion Successful",
	})
	return nil
}

// func (uh *UserHandler) AddTeam(c echo.Context) error {
// 	id := c.Param("userid")

// 	return nil
// }
