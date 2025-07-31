package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Static("/characters", "data/characters")

	e.GET("/", mainHandler)
	e.GET("/characters/", charactersHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
