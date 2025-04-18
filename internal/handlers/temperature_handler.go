package handlers

import (
	"fmt"
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/internal/viacep"
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/internal/weatherapi"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
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

	cityEscaped := url.QueryEscape(city)

	weatherResponse, err := t.WeatherAPI.GetTempByCity(city)
	if err != nil {
		panic(err)
	}
	fmt.Println(weatherResponse.Temperature)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Buscando a temperatura da cidade: " + cityEscaped))
}
