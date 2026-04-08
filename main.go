package main

import (
	"flag"
	"weather-cli/config"
	"weather-cli/service"
	"weather-cli/ui"
)

func main() {
	units    := flag.String("units", "metric", "Unit system: metric or imperial")
	forecast := flag.Bool("forecast", false, "Show 7-day forecast")

	flag.Parse()

	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}

	cities := flag.Args()

	res := service.FetchAll(cfg, cities, *units, *forecast)

	for _, r := range res {
		ui.PrintWeather(r.City, r.Weather, *units, *forecast)
	}

}
