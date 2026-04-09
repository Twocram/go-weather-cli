package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-cli/internal/config"
)

type WeatherOptions struct {
	Latitude  float64
	Longitude float64
	Units     string
	Forecast  bool
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

type DailyWeather struct {
	Time             []string  `json:"time"`
	Temperature2MMax []float64 `json:"temperature_2m_max"`
	Temperature2MMin []float64 `json:"temperature_2m_min"`
}

type WeatherResponse struct {
	Current CurrentWeather `json:"current"`
	Daily   DailyWeather   `json:"daily"`
}

func GetWeatherData(options WeatherOptions, cfg *config.Config) (*WeatherResponse, error) {
	tempUnit, windUnit := "celsius", "kmh"
	if options.Units == "imperial" {
		tempUnit, windUnit = "fahrenheit", "mph"
	}

	url := fmt.Sprintf(cfg.OpenMeteoAPIKey+"/forecast?latitude=%v&longitude=%v&current=temperature_2m,wind_speed_10m&temperature_unit=%s&wind_speed_unit=%s", options.Latitude, options.Longitude, tempUnit, windUnit)

	if options.Forecast {
		url += "&daily=temperature_2m_max,temperature_2m_min&forecast_days=7"
	} else {
		url += "&forecast_days=1"
	}

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
