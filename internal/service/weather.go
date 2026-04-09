package service

import (
	"weather-cli/internal/api"
	"weather-cli/internal/config"
)

type result struct {
	City    *api.CityInfo
	Weather *api.WeatherResponse
}

func FetchAll(cfg *config.Config, cities []string, units string, forecast bool) []result {
	results := make(chan result, len(cities))

	for _, city := range cities {
		go func(city string) {
			res, err := api.GetCityData(city, cfg)
			if err != nil {
				panic(err)
			}

			resp, err := api.GetWeatherData(api.WeatherOptions{
				Latitude:  res.Latitude,
				Longitude: res.Longitude,
				Units:     units,
				Forecast:  forecast,
			}, cfg)
			if err != nil {
				panic(err)
			}

			results <- result{
				City:    res,
				Weather: resp,
			}
		}(city)
	}

	out := make([]result, len(cities))
	for i := range out {
		out[i] = <-results
	}

	return out
}
