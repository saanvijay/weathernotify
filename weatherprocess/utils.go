package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	weathersubs "github.com/saanvijay/weathernotify/weathersubs"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func GetCurrentLocationForecast(w http.ResponseWriter, r *http.Request) {
	location, err := weathersubs.GetCurrentLocation()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	forecast, err := weathersubs.GetForeCast(location)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	jsonData, _ := json.MarshalIndent(forecast.Properties.Periods, "", " ")

	fmt.Println(string(jsonData))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	location, err := weathersubs.GetCurrentLocation()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	res := Response{
		Message: location,
		Status:  200,
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	fmt.Println(location)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func GetForecast(w http.ResponseWriter, r *http.Request) {
	latitude := r.PathValue("latitude")
	longitude := r.PathValue("longitude")
	location := fmt.Sprintf("%s, %s", latitude, longitude)

	fmt.Printf("location is %s\n", location)

	forecast, err := weathersubs.GetForeCast(location)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	jsonData, _ := json.MarshalIndent(forecast.Properties.Periods, "", " ")

	fmt.Println(string(jsonData))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
