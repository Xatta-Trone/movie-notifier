package handlers

import (
	"movie-notifier/db"
	"movie-notifier/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTracker(c echo.Context) error {
	// get the keyword

	keyword := c.FormValue("keyword")

	if keyword == "" {
		response := map[string]interface{}{"errors": "please enter the keyword"}
		return c.JSON(http.StatusBadRequest, &response)
	}

	// create instance
	tracker := entities.Tracker{Keyword: keyword}
	// save the keyword
	db.GlobalDB.Create(&tracker)

	return c.JSON(http.StatusCreated, tracker)

}

func DeleteTracker(c echo.Context) error {
	// get the keyword

	id := c.Param("id")

	if id == "" {
		response := map[string]interface{}{"errors": "please enter correct id"}
		return c.JSON(http.StatusBadRequest, &response)
	}

	// save the keyword
	db.GlobalDB.Unscoped().Delete(&entities.Tracker{}, id)

	return c.JSON(http.StatusNoContent, map[string]interface{}{})

}
