package handlers

import (
	"maatia/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CountryHandler struct {
	Service services.CountryService
}

func NewCountryHandler(s services.CountryService) *CountryHandler {
	return &CountryHandler{Service: s}
}

func (h *CountryHandler) GetCountries(c echo.Context) error {
	countries, err := h.Service.GetCountries()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get contries"})
	}

	return c.JSON(http.StatusOK, countries)
}
