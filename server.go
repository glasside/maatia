package main

import (
	"maatia/handlers"
	"maatia/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	countryService := services.NewCountryService()
	countryHanlder := handlers.NewCountryHandler(countryService)

	e.GET("/countries", countryHanlder.GetCountries)

	e.Logger.Fatal(e.Start(":8080"))
}
