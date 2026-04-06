package main

import (
	"fmt"
	"os"
	"weather-cli/service"
	"weather-cli/ui"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: weather-cli <city>")
		os.Exit(1)
	}

	cities := args[1:]

	res := service.FetchAll(cities)

	for _, r := range res {
		ui.PrintWeather(r.City, r.Weather)
	}

}
