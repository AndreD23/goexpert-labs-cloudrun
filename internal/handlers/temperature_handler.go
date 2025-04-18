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
	viaCEP     viacep.ViaCEPInterface
	weatherAPI weatherapi.WeatherAPIInterface
}

func New(weatherAPI *weatherapi.WeatherAPI) *TemperatureHandler {
	return &TemperatureHandler{
		viaCEP:     viacep.NewViaCEPService(),
		weatherAPI: weatherAPI,
	}
}

func (t *TemperatureHandler) validateZipCode(zipCode string) (string, error) {
	cleanZip := ""
	for _, char := range zipCode {
		if char >= '0' && char <= '9' {
			cleanZip += string(char)
		}
	}
	if len(cleanZip) != 8 {
		return "", fmt.Errorf("invalid zipcode: must contain exactly 8 digits")
	}
	return cleanZip, nil
}

func (t *TemperatureHandler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	zipCode := chi.URLParam(r, "zipCode")

	cleanZip, err := t.validateZipCode(zipCode)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := t.viaCEP.GetCityByZipCode(cleanZip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if city == "" {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weatherResponse, err := t.weatherAPI.GetTempByCity(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weatherResponse.Temperature)
}
