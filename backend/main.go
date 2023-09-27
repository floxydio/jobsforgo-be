package main

import (
	"jobsforgobe/database"
	"jobsforgobe/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	database.DatabaseConnect()

	routes.RouteSetup(e)

	e.Logger.Fatal(e.Start(":3000"))

}
