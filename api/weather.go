package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherOptions struct {
	Latitude  float64
	Longitude float64
}

type OpenMeteoResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	Elevation            float64 `json:"elevation"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Hourly               struct {
		Time          []string  `json:"time"`
		Temperature2M []float64 `json:"temperature_2m"`
	} `json:"hourly"`
	HourlyUnits struct {
		Temperature2M string `json:"temperature_2m"`
	} `json:"hourly_units"`
}

type CurrentWeather struct {
	Temperature2M float64 `json:"temperature_2m"`
	WindSpeed10M  float64 `json:"wind_speed_10m"`
}

type WeatherResponse struct {
	Current CurrentWeather `json:"current"`
}

func GetWeatherData(options WeatherOptions) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&current=temperature_2m,wind_speed_10m", options.Latitude, options.Longitude)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, err
	}

	return &weatherResp, nil
}
