# Weather CLI — Project Plan

## Order of implementation

- [x] Multiple cities with goroutines — accept multiple args, fetch each city concurrently with WaitGroup, collect results via channel
- [x] Units flag — `--units metric|imperial`, append unit params to API URL, display correct labels
- [ ] 7-day forecast — `--forecast` flag, request daily data, print day-by-day table
- [ ] Favorites — save/list/remove cities in `./favorites.json` (project root), fetch all favorites concurrently
