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
	path := favoritesPath

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cities []string

	if err := json.Unmarshal(data, &cities); err != nil {
		return nil, err
	}

	return cities, nil
}
