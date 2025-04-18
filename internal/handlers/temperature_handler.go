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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	city, err := viacep.GetCityByZipCode(cleanZip)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if city == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))
	}

	weatherResponse, err := t.WeatherAPI.GetTempByCity(city)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weatherResponse.Temperature)
}
