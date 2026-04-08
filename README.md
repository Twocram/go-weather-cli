# weather-cli ☁️

A tiny, cute Go CLI that asks the sky how it's feeling 🌤️

Give it one city or a few, and it will fetch the current temperature and wind speed in a clean little terminal card.

## What it does ✨

- Looks up a city with the geocoding API 🗺️
- Fetches current weather from Open-Meteo 🌦️
- Supports multiple cities at once 🏙️
- Fetches cities concurrently with goroutines ⚡
- Lets you choose `metric` or `imperial` units 🌡️
- Prints friendly styled output with `lipgloss` 🎀

## Preview 👀

```bash
go run . Moscow Tokyo Paris
```

```text
  Moscow, Russia
  Temperature  12.4°C
  Wind speed   18.0 km/h

  Tokyo, Japan
  Temperature  17.1°C
  Wind speed   9.3 km/h
```

## Quick start 🚀

### 1. Clone and install dependencies 📦

```bash
git clone <your-repo-url>
cd go-weather-cli
go mod download
```

### 2. Create a `.env` 📝

The app expects API base URLs in environment variables:

```env
OPEN_METEO_API_URL=https://api.open-meteo.com/v1
GEOCODING_API_URL=https://geocoding-api.open-meteo.com/v1
```

### 3. Run it ▶️

```bash
go run . London
```

Or with imperial units:

```bash
go run . --units imperial New York
```

Or a tiny weather parade 🎈:

```bash
go run . Seoul Berlin Lisbon
```

## Usage 🧭

```bash
go run . [--units metric|imperial] <city> [more cities...]
```

Examples:

```bash
go run . Paris
go run . --units imperial Chicago
go run . --units metric Tbilisi Yerevan Baku
```

## Docker 🐳

Build:

```bash
docker build -t weather-cli .
```

Run:

```bash
docker run --rm \
  -e OPEN_METEO_API_URL=https://api.open-meteo.com/v1 \
  -e GEOCODING_API_URL=https://geocoding-api.open-meteo.com/v1 \
  weather-cli London Rome
```

## Built with 🛠️

- Go
- Open-Meteo APIs
- `godotenv`
- `lipgloss`

## Tiny roadmap 🌱

- 7-day forecast 📅
- Favorite cities ⭐

## Notes 💡

- Right now the CLI prints current temperature and wind speed only.
- The project expects `.env` to be present when running locally.
- Errors currently fail fast with `panic`, so graceful error handling would be a nice next polish step.

## License 📄

Pick your favorite license and place it here. This little cloud deserves one ☁️
