# weather-cli ☁️

A tiny, cute Go CLI that asks the sky how it's feeling 🌤️

Give it one city or a few, and it will fetch the current weather in clean little terminal cards. It can also keep a tiny list of favorite cities for you ⭐

## What it does ✨

- Looks up a city with the geocoding API 🗺️
- Fetches current weather from Open-Meteo 🌦️
- Supports multiple cities at once 🏙️
- Fetches cities concurrently with goroutines and channels ⚡
- Lets you choose `metric` or `imperial` units 🌡️
- Shows a 7-day forecast table with `--forecast` 📅
- Saves, lists, and removes favorite cities in `favorites.json` ⭐
- Uses your saved favorites automatically when you run it without city arguments 💫
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

With forecast mode:

```bash
go run . --forecast London
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

Or ask for the next 7 days:

```bash
go run . --forecast Lisbon
```

## Usage 🧭

```bash
go run . [--units metric|imperial] [--forecast] [--action save|list|remove] <city> [more cities...]
```

Examples:

```bash
go run . Paris
go run . --units imperial Chicago
go run . --units metric Tbilisi Yerevan Baku
go run . --forecast London
go run . --action save Rome Paris
go run . --action list
go run . --action remove Rome
```

If `favorites.json` exists, you can also run:

```bash
go run .
```

That will load your saved favorite cities automatically.

## Favorites ⭐

Save a few cities:

```bash
go run . --action save Rome Paris Tokyo
```

List them:

```bash
go run . --action list
```

Remove one:

```bash
go run . --action remove Rome
```

Then later, just run:

```bash
go run .
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

- Friendlier error messages
- Duplicate-safe favorites
- Tests for forecast and favorites flows

## Notes 💡

- The project expects `.env` to be present when running locally.
- Favorites are stored in project-local `favorites.json`.
- Results are fetched concurrently, so output order may differ from the order of city arguments.
- Errors currently fail fast with `panic`, so graceful error handling would be a lovely next polish step.

## License 📄

Pick your favorite license and place it here. This little cloud deserves one ☁️
