package main

import (
	"fmt"
	"os"
	"weather-cli/api"
	"weather-cli/ui"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: weather-cli <city>")
		os.Exit(1)
	}

	cityName := args[1]

	res, err := api.GetCityData(cityName)
	if err != nil {
		panic(err)
	}

	resp, err := api.GetWeatherData(api.WeatherOptions{
		Latitude:  res.Latitude,
		Longitude: res.Longitude,
	})
	if err != nil {
		panic(err)
	}

	ui.PrintWeather(res, resp)
}
