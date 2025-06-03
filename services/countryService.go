package services

import (
	"encoding/json"
	"maatia/models"
	"net/http"
	"sort"
)

type CountryService interface {
	GetCountries() ([]models.Country, error)
}

type countryService struct{}

func NewCountryService() CountryService {
	return &countryService{}
}

func (s *countryService) GetCountries() ([]models.Country, error) {
	resp, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var raw []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var countries []models.Country

	for _, item := range raw {
		name := item["name"].(map[string]any)["official"].(string)
		var capital = ""
		if existCapital, ok := item["capital"].([]any); ok && len(existCapital) > 0 {
			if first, ok := existCapital[0].(string); ok {
				capital = first
			}
		}
		region := item["region"].(string)
		population := int(item["population"].(float64))

		countries = append(countries, models.Country{
			Name:       name,
			Capital:    capital,
			Region:     region,
			Population: population,
		})
	}

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Name < countries[j].Name
	})

	return countries, nil
}
