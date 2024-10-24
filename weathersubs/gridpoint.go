package weathersubs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GridPointResponse struct {
	Properties struct {
		Forecast string `json:"json:"forecast"`
	} `json:"Properties"`
}

func GetGridPoint(location string) (string, error) {
	url := fmt.Sprintf("https://api.weather.gov/points/%s", location)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}
	var gridPointResponse GridPointResponse
	err = json.Unmarshal(body, &gridPointResponse)
	if err != nil {
		return "", err
	}
	return gridPointResponse.Properties.Forecast, nil
}
