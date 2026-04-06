package main

import (
	"fmt"
	"os"
	"weather-cli/api"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("Please provide a city name")
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

	fmt.Println("Response:", resp)

}
