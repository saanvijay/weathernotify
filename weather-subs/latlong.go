package weathersubs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type LocationResponse struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

func getproperLocationValues(location string) string {

	location = strings.ReplaceAll(location, "-", " ")

	parts := strings.Split(location, ", ")
	var wholeNumbers []string

	for _, part := range parts {
		floatVal, err := strconv.ParseFloat(part, 64)
		if err != nil {
			fmt.Println("Error parsing float:", err)
			continue
		}

		wholeNumber := int(floatVal)
		wholeNumbers = append(wholeNumbers, strconv.Itoa(wholeNumber))
	}

	return strings.Join(wholeNumbers, ", ")

}

func GetCurrentLocation() (string, error) {
	var location LocationResponse

	locationURL := "https://ipinfo.io/json"
	response, err := http.Get(locationURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	Body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(Body, &location)
	if err != nil {
		return "", err
	}

	return getproperLocationValues(location.Loc), nil
}
