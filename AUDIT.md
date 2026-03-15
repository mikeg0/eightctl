# Eight Sleep API — Live-Validated Endpoints (2026-03-15)

> Validated with curl using OAuth2 bearer token against production API.
> Account: Pod owner, premium subscription, left side.

## Hosts

| Host | Purpose |
|------|---------|
| `auth-api.8slp.net` | OAuth2 token endpoint |
| `client-api.8slp.net` | User profiles, devices, sleep trends, intervals |
| `app-api.8slp.net` | Temperature control, routines, metrics, insights, household, and most feature endpoints |

## Authentication

```
POST https://auth-api.8slp.net/v1/tokens
Body: {"client_id":"...","client_secret":"...","grant_type":"password","username":"...","password":"..."}
Response: {"access_token":"...","token_type":"bearer","expires_in":72000,"refresh_token":"...","userId":"..."}
```

All endpoints use: `Authorization: Bearer {access_token}`

---

## AVAILABLE Endpoints (38 confirmed working)

### Auth
| Method | Host | Path | Notes |
|--------|------|------|-------|
| POST | auth-api | `/v1/tokens` | OAuth2 password grant + refresh |

### User & Profile
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | client-api | `/v1/users/me` | Current user, device list, features, timezone |
| GET | client-api | `/v1/users/{userId}` | Specific user profile |
| GET | client-api | `/v1/users/{userId}/current-device` | Side, timezone, specialization |

### Device
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | client-api | `/v1/devices/{deviceId}` | Full device state (heating levels, firmware, kelvin, wifi) |
| GET | client-api | `/v1/devices/{deviceId}?filter=...` | Filtered device fields (ownerId, leftUserId, etc.) |
| GET | client-api | `/v1/devices/{deviceId}/online` | Online status check |

### Sleep Trends (STATS — client-api)
| Method | Host | Path | Query Params | Notes |
|--------|------|------|--------------|-------|
| GET | client-api | `/v1/users/{userId}/trends` | `tz`, `from`, `to`, `include-main=false`, `include-all-sessions=true`, `model-version=v2` | **Primary daily stats**. Returns `days[]` with score, durations, sleepQualityScore, sleepRoutineScore, snoring, performanceWindows, hotFlash, sessions with full timeseries. Top-level aggregates: avgScore, avgPresenceDuration, avgSleepDuration, avgDeepPercent, avgTnt. |

### Sleep Intervals (STATS — client-api)
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | client-api | `/v1/users/{userId}/intervals` | **No session ID needed.** Returns 10 most recent intervals with full timeseries (HR, HRV, RMSSD, resp rate, bed/room temp, tnt, shortAwakes, heating), sleep stages, snoring, stageSummary. Pagination via `next` cursor. |

### Metrics (STATS — app-api)
| Method | Host | Path | Query Params | Notes |
|--------|------|------|--------------|-------|
| GET | app-api | `/v1/users/{userId}/metrics/summary` | `metrics=all` | Daily metrics: sfs, sqs, srs, sleep, light, rem, deep, hr, hrv, br, ttfa, ttgu, bedtime, waketime |
| GET | app-api | `/v1/users/{userId}/metrics/aggregate` | `metrics=all&v2=true` | Aggregated averages by period (year, etc.) |

### Insights (STATS — app-api)
| Method | Host | Path | Query Params | Notes |
|--------|------|------|--------------|-------|
| GET | app-api | `/v1/users/{userId}/insights` | `date` (optional) | Daily insights: exercise windows, performance windows, HR alerts |
| GET | app-api | `/v1/users/{userId}/llm-insights` | `from`, `to` | AI-generated sleep insights |

### Temperature
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | client-api | `/v1/users/{userId}/temperature` | Current temperature state |
| GET | app-api | `/v1/users/{userId}/temperature` | Same data, app-api host |
| GET | app-api | `/v1/users/{userId}/temperature/pod` | Pod-specific temp with `?ignoreDeviceErrors=false` |
| GET | app-api | `/v1/users/{userId}/temperature/all` | All temperature settings |
| GET | app-api | `/v1/users/{userId}/temp-events` | Temperature event history |
| GET | app-api | `/v2/smart_temperature/status/{deviceId}` | Smart temp status (left/right levels, activity) |

### Routines & Alarms
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | app-api | `/v2/users/{userId}/routines` | List all routines/alarms |
| GET | app-api | `/v2/users/{userId}/alarms` | List alarms (v2) |

### Autopilot
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | app-api | `/v1/users/{userId}/autopilotDetails` | Autopilot configuration |

### Audio
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | app-api | `/v1/audio/categories` | Audio categories |
| GET | app-api | `/v1/users/{userId}/audio/tracks` | Available audio tracks |

### Household
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | app-api | `/v1/household/users/{userId}/summary` | Household sets, devices |

### Misc
| Method | Host | Path | Notes |
|--------|------|------|-------|
| GET | app-api | `/v1/users/{userId}/release-features` | Feature flags, release dates |
| GET | app-api | `/v3/users/{userId}/subscriptions` | Subscription status |
| GET | app-api | `/v1/health-survey/test-drive` | Health survey results |
| GET | app-api | `/v1/users/{userId}/away-mode` | Away mode status |
| GET | app-api | `/v1/users/{userId}/notifications?active=true` | Active notifications |
| GET | app-api | `/v1/users/{userId}/truth-tags` | Presence truth tags |
| GET | app-api | `/v1/users/{userId}/travel/trips` | Travel trips |
| GET | app-api | `/v1/users/{userId}/challenges` | User challenges |
| GET | app-api | `/v1/users/{userId}/bedtime/recommendation` | Bedtime recommendation |
| GET | app-api | `/v1/devices/{deviceId}/priming/tasks` | Priming tasks |

---

## UNAVAILABLE Endpoints (confirmed 404)

| Host | Path | Notes |
|------|------|-------|
| client-api | `/v1/users/{userId}/preferences` | Use user profile instead |
| client-api | `/v1/users/{userId}/intervals/{sessionId}` | Session-specific lookup not supported; use `/intervals` without ID |
| client-api | `/v1/users/{userId}/metrics/summary` | **Wrong host** — works on app-api |
| client-api | `/v1/users/{userId}/metrics/aggregate` | **Wrong host** — works on app-api |
| client-api | `/v1/users/{userId}/insights` | **Wrong host** — works on app-api |
| both | `/v1/users/{userId}/sleep/heart_rate` | MCP-specific, not real |
| both | `/v1/users/{userId}/sleep/respiratory_rate` | MCP-specific, not real |
| both | `/v1/users/{userId}/sleep/timing` | MCP-specific, not real |
| both | `/v1/users/{userId}/sleep/fitness/trends` | MCP-specific, not real |
| client-api | `/v1/users/{userId}/presence` | Not available on either host |
| app-api | `/v1/users/{userId}/intervals` | **Wrong host** — works on client-api |
| app-api | `/v1/users/{userId}/trends` | **Wrong host** — works on client-api |
| app-api | `/v1/users/{userId}/audio/player` | Returns "No Associated Speaker / BaseNotPaired" |
| both | `/v1/users/{userId}/base` | app-api returns "No Associated Adjustable Base" |

---

## Key Response Schemas

### `/trends` day object keys
```
day, presenceDuration, sleepDuration, remDuration, remPercent, lightDuration,
deepDuration, deepPercent, snoreDuration, heavySnoreDuration, snorePercent,
heavySnorePercent, mitigationEvents, stoppedSnoringEvents, reducedSnoringEvents,
ineffectiveExtendedEvents, cancelledEvents, elevationDuration,
theoreticalSnorePercent, snoringReductionPercent,
elevationAutopilotAdjustmentCount, presenceStart, presenceEnd, sleepStart,
sleepEnd, tnt, mainSessionId, sessionIds, sessions, incomplete, tags,
hotFlash, performanceWindows, score, sleepQualityScore, sleepRoutineScore
```

### `sleepQualityScore` sub-keys
Each sub-metric has: `current`, `lowerRange`, `upperRange`, `lowerBound`, `upperBound`, `average`, `stdDev`, `score`, `weight`, `weighted`, `available`, `inclusive7DayAverage`

Sub-metrics: `sleepDurationSeconds`, `hrv`, `respiratoryRate`, `heartRate`, `deep`, `rem`, `waso`, `snoringDurationSeconds`, `heavySnoringDurationSeconds`, `sleepDebt`

### `sleepRoutineScore` sub-keys
Sub-metrics: `wakeupConsistency`, `sleepStartConsistency`, `bedtimeConsistency`, `latencyAsleepSeconds`, `latencyOutSeconds`

### `/intervals` object keys
```
id, deviceTimeAtUpdate, ts, stages, snoring, sleepAlgorithmVersion,
presenceAlgorithmVersion, hrvAlgorithmVersion, score, timeseries, timezone,
device, duration, stageSummary, sleepStart, sleepEnd, presenceEnd
```

### `stageSummary`
```
totalDuration, sleepDuration, outDuration, awakeDuration, lightDuration,
deepDuration, remDuration, awakeBeforeSleepDuration, awakeBetweenSleepDuration,
awakeAfterSleepDuration, outBetweenSleepDuration, wasoDuration,
deepPercentOfSleep, remPercentOfSleep, lightPercentOfSleep
```

### `timeseries` keys
```
tnt, tempRoomC, tempBedC, respiratoryRate, nemeanRespiratoryRate,
nemeanRespiratoryRateNightly, heartRate, heating, hrv, rmssd, shortAwakes
```

### `/metrics/summary` metric names
```
sfs (sleep fitness score), avg_sfs, sqs (sleep quality score), avg_sqs,
srs (sleep routine score), avg_srs, sleep (duration secs), sds (sleep duration score),
light, rem, rem_percent, deep, deep_percent, hr, hrv, br (breathing rate),
ttfa (time to fall asleep), ttgu (time to get up), bedtime, waketime
```

---

## Host Routing Rule

| Endpoint category | Host |
|-------------------|------|
| Auth (tokens) | `auth-api.8slp.net` |
| User profile, device info, trends, intervals | `client-api.8slp.net` |
| Everything else (temperature, metrics, insights, routines, household, etc.) | `app-api.8slp.net` |
