package main

import (
	"fmt"
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/configs"
	"net/http"
)

func main() {
	config := configs.NewConfig()

	fmt.Println("Olá")
	fmt.Println(config.WeatherAPI)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World, estou no cloud run!"))
	})
	http.ListenAndServe(":8080", nil)
}
