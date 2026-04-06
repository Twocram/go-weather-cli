# Weather CLI — Project Plan

## Order of implementation

- [ ] Multiple cities with goroutines — accept multiple args, fetch each city concurrently with WaitGroup, collect results via channel
- [ ] Units flag — `--units metric|imperial`, append unit params to API URL, display correct labels
- [ ] 7-day forecast — `--forecast` flag, request daily data, print day-by-day table
- [ ] Favorites — save/list/remove cities in `~/.weather-cli/favorites.json`, fetch all favorites concurrently
