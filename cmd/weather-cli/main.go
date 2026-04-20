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
		ui.PrintError(err)
		os.Exit(1)
	}

	switch *action {
	case "save":
		if len(flag.Args()) == 0 {
			ui.PrintError(fmt.Errorf("usage: weather-cli --action save <city>..."))
			os.Exit(1)
		}

		if err := service.SaveFavorites(flag.Args()); err != nil {
			ui.PrintError(err)
			os.Exit(1)
		}

		return
	case "list":
		cities, err := service.LoadFavorites()
		if err != nil {
			ui.PrintError(err)
			os.Exit(1)
		}
		for _, city := range cities {
			fmt.Println(city)
		}
		return
	case "remove":
		if len(flag.Args()) == 0 {
			ui.PrintError(fmt.Errorf("usage: weather-cli --action save <city>..."))
			os.Exit(1)
		}
		if err := service.RemoveFavorites(flag.Args()); err != nil {
			ui.PrintError(err)
			os.Exit(1)
		}
		return
	}

	cities := flag.Args()

	if len(cities) == 0 {
		if service.FavoritesFileExist() {
			loaded, err := service.LoadFavorites()
			if err != nil {
				ui.PrintError(err)
				os.Exit(1)
			}
			cities = loaded
		} else {
			ui.PrintError(fmt.Errorf("usage: weather-cli [--units metric|imperial] <city>..."))
			os.Exit(1)
		}
	}

	res := service.FetchAll(cfg, cities, *units, *forecast)

	for _, r := range res {
		ui.PrintWeather(r.City, r.Weather, *units, *forecast)
	}

}
