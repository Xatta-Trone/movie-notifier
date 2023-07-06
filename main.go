package main

import (
	"log"
	"movie-notifier/db"
	"movie-notifier/handlers"
	"net/http"
	"os"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)


func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init db
	database := db.InitDB()
	// Migrate the schema
	db.MigrateDb(database)

	// init http server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/trackers",handlers.CreateTracker)
	e.DELETE("/trackers/:id",handlers.DeleteTracker)

	// windows fix
	URL := ""
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "1323"
	}

	if runtime.GOOS == "windows" {
		URL = "localhost:" + PORT
	} else {
		URL = ":" + PORT
	}

	e.Logger.Fatal(e.Start(URL))
}


