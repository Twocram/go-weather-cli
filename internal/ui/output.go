package ui

import (
	"fmt"
	"strings"
	"weather-cli/internal/api"

	"charm.land/lipgloss/v2"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 2)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	valueStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA"))

	containerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 3)
)

func PrintWeather(city *api.CityInfo, weather *api.WeatherResponse, units string, forecast bool) {
	tempChar, windChar := "°C", "km/h"

	if units == "imperial" {
		tempChar, windChar = "°F", "mph"
	}

	title := titleStyle.Render(fmt.Sprintf("  %s, %s", city.Name, city.Country))

	temp := fmt.Sprintf("%s  %s",
		labelStyle.Render("Temperature"),
		valueStyle.Render(fmt.Sprintf("%.1f%s", weather.Current.Temperature2M, tempChar)),
	)
	wind := fmt.Sprintf("%s  %s",
		labelStyle.Render("Wind speed "),
		valueStyle.Render(fmt.Sprintf("%.1f %s", weather.Current.WindSpeed10M, windChar)),
	)

	body := containerStyle.Render(fmt.Sprintf("%s\n%s", temp, wind))

	fmt.Println()
	fmt.Println(title)
	fmt.Println(body)

	if forecast {
		highStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF6B6B"))
		lowStyle  := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#74C7EC"))
		headStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4"))

		dateW := lipgloss.NewStyle().Width(14)
		highW := lipgloss.NewStyle().Width(12)
		lowW  := lipgloss.NewStyle().Width(12)

		header := dateW.Render(headStyle.Render("Date")) +
			highW.Render(headStyle.Render("High")) +
			lowW.Render(headStyle.Render("Low"))

		sep := labelStyle.Render(strings.Repeat("─", 36))

		rows := []string{header, sep}
		for i, day := range weather.Daily.Time {
			row := dateW.Render(labelStyle.Render(day)) +
				highW.Render(highStyle.Render(fmt.Sprintf("↑ %.1f%s", weather.Daily.Temperature2MMax[i], tempChar))) +
				lowW.Render(lowStyle.Render(fmt.Sprintf("↓ %.1f%s", weather.Daily.Temperature2MMin[i], tempChar)))
			rows = append(rows, row)
		}

		fmt.Println(containerStyle.Render(strings.Join(rows, "\n")))
	}
}
