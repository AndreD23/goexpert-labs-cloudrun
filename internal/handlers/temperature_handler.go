package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/AndreD23/goexpert-labs-cloudrun/internal/viacep"
	"github.com/AndreD23/goexpert-labs-cloudrun/internal/weatherapi"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type TemperatureHandler struct {
	*weatherapi.WeatherAPI
}

func New(weatherAPI *weatherapi.WeatherAPI) *TemperatureHandler {
	return &TemperatureHandler{
		WeatherAPI: weatherAPI,
	}
}

func (t *TemperatureHandler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	zipCode := chi.URLParam(r, "zipCode")
	fmt.Println(zipCode)

	city, err := viacep.GetCityByZipCode(zipCode)
	if err != nil {
		panic(err)
	}

	weatherResponse, err := t.WeatherAPI.GetTempByCity(city)
	if err != nil {
		panic(err)
	}
	fmt.Println(weatherResponse.Temperature)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weatherResponse.Temperature)
}
