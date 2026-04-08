package service

import (
	"encoding/json"
	"errors"
	"os"
)

const favoritesPath = "favorites.json"

func FavoritesFileExist() bool {
	_, err := os.Stat(favoritesPath)
	return !errors.Is(err, os.ErrNotExist)
}

func LoadFavorites() ([]string, error) {
	data, err := os.ReadFile(favoritesPath)
	if err != nil {
		return nil, err
	}

	var cities []string

	if err := json.Unmarshal(data, &cities); err != nil {
		return nil, err
	}

	return cities, nil
}

func SaveFavorites(newCities []string) error {
	data, err := os.ReadFile(favoritesPath)
	var cities []string

	if err == nil {
		json.Unmarshal(data, &cities)
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	for _, city := range newCities {
		cities = append(cities, city)
	}

	data, err = json.Marshal(cities)
	if err != nil {
		return err
	}

	return os.WriteFile(favoritesPath, data, 0644)
}

func RemoveFavorites(toRemove []string) error {
	data, err := os.ReadFile(favoritesPath)
	if err != nil {
		return err
	}

	var cities []string
	if err := json.Unmarshal(data, &cities); err != nil {
		return err
	}

	removeSet := make(map[string]bool, len(toRemove))
	for _, c := range toRemove {
		removeSet[c] = true
	}

	filtered := cities[:0]
	for _, c := range cities {
		if !removeSet[c] {
			filtered = append(filtered, c)
		}
	}

	data, err = json.Marshal(filtered)
	if err != nil {
		return err
	}

	return os.WriteFile(favoritesPath, data, 0644)
}
