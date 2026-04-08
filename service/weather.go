package service

import (
	"sync"
	"weather-cli/api"
	"weather-cli/config"
)

type result struct {
	City    *api.CityInfo
	Weather *api.WeatherResponse
}

func FetchAll(cfg *config.Config, cities []string, units string) []result {
	var wg sync.WaitGroup

	results := make([]result, len(cities))

	for i, city := range cities {
		wg.Add(1)
		go func(i int, city string) {
			defer wg.Done()

			res, err := api.GetCityData(city, cfg)
			if err != nil {
				panic(err)
			}

			resp, err := api.GetWeatherData(api.WeatherOptions{
				Latitude:  res.Latitude,
				Longitude: res.Longitude,
				Units:     units,
			}, cfg)

			if err != nil {
				panic(err)
			}

			results[i] = result{res, resp}
		}(i, city)
	}

	wg.Wait()

	return results
}
