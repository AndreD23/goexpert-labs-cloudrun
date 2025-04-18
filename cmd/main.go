package main

import (
	"fmt"
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/configs"
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/internal/viacep"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	config := configs.NewConfig()

	fmt.Println("Olá")
	fmt.Println(config.WeatherAPI)

	r := chi.NewRouter()
	r.Get("/", HelloHandler)
	r.Get("/{zipCode}", TemperatureHandler)
	http.ListenAndServe(":8080", r)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World, estou no cloud run!"))
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	zipCode := chi.URLParam(r, "zipCode")
	fmt.Println(zipCode)

	city, err := viacep.GetCityByZipCode(zipCode)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Buscando a temperatura da cidade: " + city))
}
