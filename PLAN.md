# Weather CLI — Project Plan

## Shipped

- [x] Multiple cities support
  Accept multiple positional city arguments and fetch them concurrently.
- [x] Units flag
  Support `--units metric|imperial` and display the correct temperature and wind labels.
- [x] 7-day forecast
  Support `--forecast` and print a daily high/low table.
- [x] Favorites workflow
  Save, list, and remove cities in `./favorites.json`.
- [x] Favorites as default input
  When no city arguments are passed, load favorites automatically if `favorites.json` exists.
- [x] Channel-based fetch flow
  `FetchAll` now fans out requests with goroutines and collects results through a results channel.

## Next up

- [ ] Preserve input order in concurrent results
  Current channel collection returns cities in completion order, not argument order.
- [ ] Graceful error handling
  Replace `panic` paths with friendly CLI errors and partial-result handling.
- [ ] Better favorites UX
  Avoid duplicate saved cities and print confirmations for save/remove/list actions.
- [ ] Config polish
  Clarify naming around API base URLs and improve `.env` setup messaging.
- [ ] Tests
  Add coverage for favorites, forecast mode, and no-argument fallback behavior.
