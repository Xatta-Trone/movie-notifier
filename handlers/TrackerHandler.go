package handlers

import (
	"movie-notifier/db"
	"movie-notifier/entities"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetTrackers(c echo.Context) error {
	// get the keyword

	page := c.QueryParam("page")

	if page == "" {
		page = "1"
	}

	pageInt, _ := strconv.Atoi(page)

	perPage := c.QueryParam("per_page")

	if perPage == "" {
		perPage = "10"
	}

	perPageInt, _ := strconv.Atoi(perPage)

	offset := (pageInt - 1) * perPageInt

	// get the data 

	trackers := []entities.Tracker{}
	
	// save the keyword
	db.GlobalDB.Order("id desc").Offset(offset).Limit(perPageInt).Find(&trackers)

	return c.JSON(http.StatusOK, &trackers)

}

func CreateTracker(c echo.Context) error {
	// get the keyword

	keyword := c.FormValue("keyword")

	if keyword == "" {
		response := map[string]interface{}{"errors": "please enter the keyword"}
		return c.JSON(http.StatusBadRequest, &response)
	}

	keyword = strings.ToLower(keyword)

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
