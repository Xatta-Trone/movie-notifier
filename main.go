package main

import (
	"log"
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

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

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
