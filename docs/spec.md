# eightctl Specification

## Purpose
Eight Sleep Pod control + data-export CLI, written in Go. Targets macOS/Linux users who want a terminal tool for pod automations, metrics export, and feature toggles that the mobile app exposes but the vendor does not document.

## Reality of the API
- Eight Sleep does **not** publish a stable public API; we rely on the same cloud endpoints the mobile apps use.
- Three API hosts:
  - `auth-api.8slp.net` — OAuth2 token endpoint
  - `client-api.8slp.net` — user profiles, devices, sleep trends, intervals
  - `app-api.8slp.net` — temperature, metrics, insights, routines, household, and most feature endpoints
- Default OAuth client creds extracted from Android APK 7.39.17:
  - `client_id`: `0894c7f33bb94800a03f1f4df13a4f38`
  - `client_secret`: `f0954a3ed5763ba3d06834c73731a32f15f168f47d4f164751275def86db0c76`
- Auth flow: password grant at `https://auth-api.8slp.net/v1/tokens`; fallback legacy `/login` session token.
- Throttling: 429s observed; client retries with small delay and re-auths on 401.

## Configuration & Auth
- Config file: `~/.config/eightctl/config.yaml`; env prefix `EIGHTCTL_`; flags override env override file.
- Fields: `email`, `password`, optional `user_id`, `client_id`, `client_secret`, `timezone`, `output`, `fields`, `verbose`.
- Permissions check warns if config is more permissive than `0600`.

## CLI Surface (implemented & validated)

Core: `on`, `off`, `temp <level>`, `status`, `whoami`, `version`, `logout`.

Schedules & daemon:
- `schedule list|create|update|delete|next` (cloud temperature schedules via app-api)
- `daemon` (YAML-based scheduler with PID guard, dry-run, timezone override)

Alarms:
- `alarm list` (routines via app-api v2)
- `alarm snooze|dismiss|dismiss-all|vibration-test`

Temperature modes:
- `tempmode nap on|off|extend|status` (app-api)
- `tempmode hotflash on|off|status` (app-api)
- `tempmode events --from --to` (temperature event history, app-api)

Audio:
- `audio tracks|categories` (app-api)

Adjustable base:
- `base info|angle` (app-api; requires paired adjustable base)

Device:
- `device info|online|priming-tasks`

Metrics & insights (all validated):
- `metrics trends --from --to` (client-api)
- `metrics intervals [--cursor]` (client-api — no session ID needed)
- `metrics summary` (app-api)
- `metrics aggregate` (app-api)
- `metrics insights [--date]` (app-api)
- `metrics llm-insights --from --to` (app-api)
- `sleep day --date`, `sleep range --from --to` (client-api trends)

Autopilot:
- `autopilot details` (app-api)

Travel:
- `travel trips` (app-api)

Household:
- `household summary` (app-api)

Presence:
- `presence` (derived from device polling)

## Output & UX
- Output formats: table (default), json, csv via `--output`; `--fields` to select columns.
- Logs via charmbracelet/log; `--verbose` for debug; `--quiet` hides config notice.

## Daemon Behavior
- Reads YAML schedule (time, action on|off|temp, temperature with unit), minute tick, executes once per day, PID guard, SIGINT/SIGTERM graceful stop.

## Testing & Quality Gates
- `go test ./...` (fast compile checks) — run before handoff.
- Live checks: `eightctl status`, `metrics summary`, `metrics intervals`, `sleep day` with test creds.

## Prior Work (references)
- Go CLI `clim8`: https://github.com/blacktop/clim8
- MCP server (Node/TS): https://github.com/elizabethtrykin/8sleep-mcp
- Python library `pyEight`: https://github.com/mezz64/pyEight
- Home Assistant integrations: https://github.com/lukas-clarke/eight_sleep and https://github.com/grantnedwards/eight-sleep
- Homebridge plugin: https://github.com/nfarina/homebridge-eightsleep
- Additional notes on API stability: https://www.reddit.com/r/EightSleep/comments/15ybfrv/eight_sleep_removed_smart_home_capabilities/
