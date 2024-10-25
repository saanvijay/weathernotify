package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	Host := "localhost"
	Port := "8080"

	mux := http.NewServeMux()
	mux.HandleFunc("GET /getlocation", GetLocation)
	mux.HandleFunc("GET /getforecast/{latitude}/{longitude}", GetForecast)
	mux.HandleFunc("GET /getcurrentlocationforecast", GetCurrentLocationForecast)

	// run the forcast for every 15 mins
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()
	// Run the task immediately on start, then every 15 minutes
	kafkaProduceForcast()

	for range ticker.C {
		kafkaProduceForcast()
	}

	endpoint := fmt.Sprintf("%s:%s", Host, Port)
	fmt.Printf("Server listening %s\n", endpoint)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
		return
	}

}
