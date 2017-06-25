package main

import (
	"log"
	"net/http"

	"github.com/crucial7/api-weather/handlers"
)

func main() {

	http.HandleFunc("/weatherByLatLongAndDate/", handlers.WeatherByLatLongAndDate)
	log.Println("Server listening on port 3002")
	log.Println("\tRoutes:")
	log.Println("\t\tGET /weatherByLatLongAndDate")
	log.Fatal(http.ListenAndServe(":3002", nil))
}
