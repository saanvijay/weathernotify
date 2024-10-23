package main

import (
	"encoding/json"
	"fmt"

	weathersubs "github.com/saanvijay/weathernotify/weather-subs"
)

func main() {

	location, err := weathersubs.GetCurrentLocation()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Println(location)

	grid, err := weathersubs.GetGridPoint(location)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Println(string(grid))

	forecast, err := weathersubs.GetForeCast(location)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	jsonData, _ := json.MarshalIndent(forecast.Properties.Periods, "", " ")

	fmt.Println(string(jsonData))

}
