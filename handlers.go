package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func mainHandler(c echo.Context) error {
	f, err := os.ReadFile("./static/index.html")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.HTML(http.StatusOK, string(f))
}

func charactersHandler(c echo.Context) error {
	data, err := readCharacters()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSONPretty(http.StatusOK, data, "  ")
}
