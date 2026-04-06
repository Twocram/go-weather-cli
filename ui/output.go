package ui

import (
	"fmt"
	"weather-cli/api"

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

func PrintWeather(city *api.CityInfo, weather *api.WeatherResponse) {
	title := titleStyle.Render(fmt.Sprintf("  %s, %s", city.Name, city.Country))

	temp := fmt.Sprintf("%s  %s",
		labelStyle.Render("Temperature"),
		valueStyle.Render(fmt.Sprintf("%.1f°C", weather.Current.Temperature2M)),
	)
	wind := fmt.Sprintf("%s  %s",
		labelStyle.Render("Wind speed "),
		valueStyle.Render(fmt.Sprintf("%.1f km/h", weather.Current.WindSpeed10M)),
	)

	body := containerStyle.Render(fmt.Sprintf("%s\n%s", temp, wind))

	fmt.Println()
	fmt.Println(title)
	fmt.Println(body)
}
