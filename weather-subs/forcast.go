package weathersubs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ForecastResponse struct {
	Properties struct {
		Periods []struct {
			Name          string `json:"name"`
			Temperature   int    `json:"temperature"`
			WindSpeed     string `json:"windSpeed"`
			ShortForecast string `json:"shortForecast"`
		} `json:periods`
	} `json:Properties`
}

func GetForeCast(location string) (*ForecastResponse, error) {

	var forecastResponse ForecastResponse
	forecastURL := fmt.Sprintf("https://api.weather.gov/gridpoints/LWX/%s/forecast", location)

	response, err := http.Get(forecastURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &forecastResponse)
	if err != nil {
		return nil, err
	}
	return &forecastResponse, nil
}
