package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenMeteoAPIKey string
	GeoCodingAPIKey string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	openMeteoUrl := os.Getenv("OPEN_METEO_API_URL")

	if openMeteoUrl == "" {
		return nil, fmt.Errorf("OPEN_METEO_API_URL is not set")
	}

	geoCodingUrl := os.Getenv("GEOCODING_API_URL")

	if geoCodingUrl == "" {
		return nil, fmt.Errorf("GEOCODING_API_URL is not set")
	}

	return &Config{
		OpenMeteoAPIKey: openMeteoUrl,
		GeoCodingAPIKey: geoCodingUrl,
	}, nil
}
