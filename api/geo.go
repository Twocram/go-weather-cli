package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CityInfo struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Elevation   float64 `json:"elevation"`
	FeatureCode string  `json:"feature_code"`
	CountryCode string  `json:"country_code"`
	Admin1ID    int     `json:"admin1_id"`
	Admin2ID    int     `json:"admin2_id"`
	Timezone    string  `json:"timezone"`
	Population  int     `json:"population"`
	CountryID   int     `json:"country_id"`
	Country     string  `json:"country"`
	Admin1      string  `json:"admin1"`
	Admin2      string  `json:"admin2"`
}

type geoResponse struct {
	Results []CityInfo `json:"results"`
}

func GetCityData(city string) (*CityInfo, error) {
	response, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?name=" + city + "&count=1")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var geoResp geoResponse
	if err := json.NewDecoder(response.Body).Decode(&geoResp); err != nil {
		return nil, err
	}

	if len(geoResp.Results) == 0 {
		return nil, fmt.Errorf("city not found: %s", city)
	}

	return &geoResp.Results[0], nil
}
