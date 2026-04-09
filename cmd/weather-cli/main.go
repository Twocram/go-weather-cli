package main

import (
	"flag"
	"fmt"
	"os"
	"weather-cli/internal/config"
	"weather-cli/internal/service"
	"weather-cli/internal/ui"
)

func main() {
	units := flag.String("units", "metric", "Unit system: metric or imperial")
	forecast := flag.Bool("forecast", false, "Show 7-day forecast")
	action := flag.String("action", "none", "Action: save|list|remove")

	flag.Parse()

	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}

	switch *action {
	case "save":
		if len(flag.Args()) == 0 {
			fmt.Println("Usage: weather-cli --action save <city>...")
			os.Exit(1)
		}

		if err := service.SaveFavorites(flag.Args()); err != nil {
			panic(err)
		}

		return
	case "list":
		cities, err := service.LoadFavorites()
		if err != nil {
			panic(err)
		}
		for _, city := range cities {
			fmt.Println(city)
		}
		return
	case "remove":
		if len(flag.Args()) == 0 {
			fmt.Println("Usage: weather-cli --action save <city>...")
			os.Exit(1)
		}
		service.RemoveFavorites(flag.Args())
		return
	}

	cities := flag.Args()

	if len(cities) == 0 {
		if service.FavoritesFileExist() {
			loaded, err := service.LoadFavorites()
			if err != nil {
				panic(err)
			}
			cities = loaded
		} else {
			fmt.Println("Usage: weather-cli [--units metric|imperial] <city>...")
			os.Exit(1)
		}
	}

	res := service.FetchAll(cfg, cities, *units, *forecast)

	for _, r := range res {
		ui.PrintWeather(r.City, r.Weather, *units, *forecast)
	}

}
