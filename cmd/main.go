package main

import (
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/configs"
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/internal/handlers"
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/internal/weatherapi"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	config := configs.NewConfig()
	weather := weatherapi.NewWeatherAPI(config.WeatherAPIKey)
	temperatureHandler := handlers.New(weather)

	r := chi.NewRouter()
	r.Get("/", HelloHandler)
	r.Get("/{zipCode}", temperatureHandler.GetTemperature)
	http.ListenAndServe(":8080", r)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World, estou no cloud run!"))
}
