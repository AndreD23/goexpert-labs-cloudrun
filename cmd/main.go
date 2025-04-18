package main

import (
	"github.com/AndreD23/goexpert-labs-cloudrun/configs"
	"github.com/AndreD23/goexpert-labs-cloudrun/internal/handlers"
	"github.com/AndreD23/goexpert-labs-cloudrun/internal/weatherapi"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	config := configs.NewConfig()
	weather := weatherapi.NewWeatherAPI(config.WeatherAPIKey)
	temperatureHandler := handlers.New(weather)

	r := chi.NewRouter()
	r.Get("/{zipCode}", temperatureHandler.GetTemperature)
	http.ListenAndServe(":8080", r)
}
