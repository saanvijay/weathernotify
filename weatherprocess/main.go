package main

import (
	"fmt"
	"net/http"
)

func main() {

	Host := "localhost"
	Port := "8080"

	mux := http.NewServeMux()
	mux.HandleFunc("GET /getlocation", GetLocation)
	mux.HandleFunc("GET /getforecast/{latitude}/{longitude}", GetForecast)
	mux.HandleFunc("GET /getcurrentlocationforecast", GetCurrentLocationForecast)

	endpoint := fmt.Sprintf("%s:%s", Host, Port)
	fmt.Printf("Server listening %s\n", endpoint)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
		return
	}

}
