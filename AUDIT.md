# Eight Sleep API Endpoint Documentation (audit)

> **Source:** Decompiled from `eightsleep-base.apk` via `jadx`.
> **Status:** Live-validated with `curl` on 2026-03-14. `available` means the exact method/path responded from the API; `unavailable` means the API returned `Cannot METHOD /path`. Mutating routes were probed with fake IDs and invalid/empty JSON bodies to avoid changing account state.

---

## Configuration

### Hosts

| Environment | Base URL |
|-------------|----------|
| Production | `https://{subdomain}.8slp.net/` |
| Staging | `https://{subdomain}.staging.8slp.net/` |

**Subdomains:**
- Auth: `auth-api.8slp.net`
- API: `client-api.8slp.net`

### Embedded Client Credentials (from APK)

```
client_id:     0894c7f33bb94800a03f1f4df13a4f38
client_secret: f0954a3ed5763ba3d06834c73731a32f15f168f47d4f164751275def86db0c76
grant_type:    password
```

### Common Headers

```
Authorization: Bearer {access_token}
Content-Type:  application/json
```

---

## Required IDs

Each endpoint requires one or more IDs resolved at runtime:

| ID | How to obtain |
|----|---------------|
| `userId` | Returned in login response (`userId` field) |
| `deviceId` | From device list or user profile |
| `householdId` | From `GET v1/household/users/{userId}/summary` |
| `tripId` / `planId` | From travel trip endpoints |
| `sessionId` | From sleep session/interval endpoints |
| `alarmId` | From alarm list endpoint |

---

## Authentication

**Host:** `auth-api.8slp.net`
**Source:** `jadx-out/sources/C9/a.java`

### Login

```
POST v1/tokens
```

Status: `available`

Body:

```json
{
  "client_id": "...",
  "client_secret": "...",
  "grant_type": "password",
  "username": "user@example.com",
  "password": "..."
}
```

Response:

```json
{
  "access_token": "...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "...",
  "userId": "..."
}
```

### Refresh Token

```
POST v1/tokens
```

Status: `available`

Body:

```json
{
  "client_id": "...",
  "client_secret": "...",
  "grant_type": "refresh_token",
  "refresh_token": "..."
}
```

### Register User

```
POST v1/users
```

Status: `unavailable`

Body: `RegisterNewUserRequest`

### Reset Password

```
POST v1/users/password-temporary
```

Status: `unavailable`

Body: `ResetPasswordRequest`

---

## Devices

**Host:** `client-api.8slp.net`
**Source:** `jadx-out/sources/Ia/InterfaceC1265o.java`

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v1/devices/{deviceId}` | Get device details | — | available |
| GET | `v1/devices/{deviceId}/online` | Check device online status | — | available |
| PUT | `v1/devices/{deviceId}` | Update device | `UpdateDeviceRequest` | available |
| PUT | `v1/devices/{deviceId}/owner` | Set device owner | `SetDeviceOwnerRequest` | available |
| PUT | `v1/devices/{deviceId}/peripherals` | Update peripherals | `UpdateDevicePeripheralRequest` | available |
| PATCH | `v1/devices/{deviceId}/peripherals` | Patch peripherals | `SetDevicePeripheralsRequest` | available |

### Device Priming & Pairing

**Source:** `jadx-out/sources/Ia/InterfaceC1266p.java`

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| POST | `v1/devices/{deviceId}/priming/tasks` | Create priming task | `PrimingTaskRequest` | unavailable |
| GET | `v1/devices/{deviceId}/priming/tasks` | Get priming tasks | — | unavailable |
| DELETE | `v1/devices/{deviceId}/priming/tasks` | Delete priming task | — | unavailable |
| GET | `v1/devices/{deviceId}/priming/schedule` | Get priming schedule | — | unavailable |
| PUT | `v1/devices/{deviceId}/priming/schedule` | Update priming schedule | `PrimingScheduleNetwork` | unavailable |
| GET | `v1/devices/{deviceId}/warranty` | Get warranty info | — | unavailable |
| POST | `v1/devices/{deviceId}/auto-pairing/start` | Start auto pairing | `HubAutoPairingRequest` | unavailable |
| GET | `v1/devices/{deviceId}/auto-pairing/status/{pairingId}` | Check pairing status | — | unavailable |
| POST | `v1/devices/{deviceId}/security/key` | Get device key | `GetDeviceKeyRequest` | unavailable |
| POST | `v2/devices/{deviceId}/vibration-test` | Start vibration test | `DeviceVibrationTestRequest` | unavailable |
| POST | `v2/devices/{deviceId}/vibration-test/stop` | Stop vibration test | — | unavailable |

---

## Users & Profile

**Host:** `client-api.8slp.net`

| Method | Path | Purpose | Status |
|--------|------|---------|--------|
| GET | `v1/users/me` | Bootstrap — get current user (inferred) | available |

---

## Temperature

**Source:** `jadx-out/sources/dc/a.java`, `dc/b.java`

### Temperature Settings

| Method | Path | Purpose | Query Params | Body | Status |
|--------|------|---------|--------------|------|--------|
| GET | `v1/users/{userId}/temperature/all` | Get all temperature settings | — | — | unavailable |
| PUT | `v1/users/{userId}/temperature/{deviceType}` | Update temperature settings | `ignoreDeviceErrors` | `TemperatureSettingsRequest` | unavailable |
| GET | `v1/users/{userId}/temp-events` | Get temperature events | — | — | unavailable |

### Nap Mode

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v1/users/{userId}/temperature/nap-mode` | Get nap mode settings | — | unavailable |
| GET | `v1/users/{userId}/temperature/nap-mode/status` | Get nap mode status | — | unavailable |
| POST | `v1/users/{userId}/temperature/nap-mode/activate` | Activate nap mode | `StartNapRequest` | unavailable |
| POST | `v1/users/{userId}/temperature/nap-mode/extend` | Extend nap mode | `ExtendNapRequest` | unavailable |
| POST | `v1/users/{userId}/temperature/nap-mode/deactivate` | Deactivate nap mode | — | unavailable |

### Hot Flash Mode

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v1/users/{userId}/temperature/hot-flash-mode` | Get hot flash settings | — | unavailable |
| POST | `v1/users/{userId}/temperature/hot-flash-mode` | Update hot flash mode | `HotFlashModeSettingsRequest` | unavailable |
| POST | `v1/users/{userId}/temperature/hot-flash-mode/activate` | Activate hot flash mode | — | unavailable |
| POST | `v1/users/{userId}/temperature/hot-flash-mode/deactivate` | Deactivate hot flash mode | — | unavailable |
| DELETE | `v1/users/{userId}/temperature/hot-flash-mode` | Delete hot flash mode | — | unavailable |

---

## Alarms

**Source:** `jadx-out/sources/J9/InterfaceC5338a.java`

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v2/users/{userId}/alarms` | List alarms | — | unavailable |
| POST | `v1/users/{userId}/alarms` | Create alarm | `CreateAlarmRequest` | unavailable |
| PUT | `v1/users/{userId}/alarms/{alarmId}` | Update alarm | `UpdateAlarmRequest` | unavailable |
| DELETE | `v1/users/{userId}/alarms/{alarmId}` | Delete alarm | — | unavailable |
| POST | `v1/users/{userId}/alarms/{alarmId}/snooze` | Snooze alarm | `SnoozeAlarmRequest` | unavailable |
| POST | `v1/users/{userId}/alarms/{alarmId}/dismiss` | Dismiss alarm | `IgnoreDeviceErrorsRequest` | unavailable |
| POST | `v1/users/{userId}/alarms/active/dismiss-all` | Dismiss all active alarms | `IgnoreDeviceErrorsRequest` | unavailable |
| POST | `v1/users/{userId}/vibration-test` | Run vibration test | `VibrationTestRequest` | unavailable |
| GET | `v1/users/{userId}/temporary-mode/nap-mode` | Get nap mode alarm settings | — | unavailable |
| PUT | `v1/users/{userId}/temporary-mode/nap-mode` | Update nap mode alarm settings | `NapModeAlarmSettingsNetwork` | unavailable |

---

## Sleep & Metrics

**Source:** `jadx-out/sources/ob/a.java`, `ob/b.java`

### Metrics

| Method | Path | Purpose | Query Params | Status |
|--------|------|---------|--------------|--------|
| GET | `v1/users/{userId}/metrics/summary` | Get metrics summary | `from` (date), `to` (date), `tz` (timezone), `metrics` | unavailable |
| GET | `v1/users/{userId}/metrics/aggregate` | Get metrics aggregate | `to`, `tz`, `metrics`, `periods`, `refreshCache`, `v2=true` | unavailable |

### Trends & Sessions

| Method | Path | Purpose | Query Params | Body | Status |
|--------|------|---------|--------------|------|--------|
| GET | `v1/users/{userId}/trends` | Get sleep trends | `tz`, `from`, `to`, `include-main`, `include-all-sessions`, `model-version`, `consistent-read` | — | available |
| POST | `v1/users/{userId}/intervals/{sessionId}` | Update sleep session | — | `UpdateSleepSessionRequest` | unavailable |
| PUT | `v1/users/{userId}/intervals/{sessionId}` | Stop sleep session | — | `StopSleepSessionRequest` | available |
| DELETE | `v1/users/{userId}/intervals/{sessionId}` | Delete sleep session | — | — | available |
| POST | `v1/users/{userId}/feedback` | Post metrics feedback | — | `MetricsFeedbackRequest` | available |

---

## Autopilot

**Source:** `jadx-out/sources/k9/a.java`

| Method | Path | Purpose | Query Params | Body | Status |
|--------|------|---------|--------------|------|--------|
| GET | `v1/users/{userId}/autopilotDetails` | Get autopilot details | — | — | unavailable |
| GET | `v1/users/{userId}/autopilotDetails/autopilotRecap` | Get autopilot recap | `day` (date), `tz` (timezone) | — | unavailable |
| GET | `v1/users/{userId}/autopilot-history` | Get autopilot history | — | — | unavailable |
| PUT | `v1/users/{userId}/autopilotDetails/snoringMitigation` | Update snoring mitigation | — | `SnoringMitigationNetworkModelWrapper` | unavailable |
| GET | `v1/users/{userId}/level-suggestions-mode` | Get level suggestions mode | — | — | unavailable |
| PUT | `v1/users/{userId}/level-suggestions-mode` | Update level suggestions | — | `AutopilotSettingsRequest` | unavailable |

---

## Adjustable Base

**Source:** `jadx-out/sources/P8/a.java`

| Method | Path | Purpose | Query Params | Body | Status |
|--------|------|---------|--------------|------|--------|
| GET | `v1/users/{userId}/base` | Get base state | — | — | unavailable |
| GET | `v2/users/{userId}/base/presets` | Get base presets (v2) | — | — | unavailable |
| GET | `v1/users/{userId}/base/presets` | Get base presets | — | — | unavailable |
| POST | `v1/users/{userId}/base/presets` | Create base preset | — | `UpdateBaseDefaultPresetRequest` | unavailable |
| POST | `v1/users/{userId}/base/angle` | Set base angle | `ignoreDeviceErrors` | `SetBaseAngleRequest` | unavailable |
| DELETE | `v1/users/{userId}/base/angle` | Delete base angle | `ignoreDeviceErrors` | — | unavailable |
| DELETE | `v1/devices/{deviceId}/base` | Remove device base | — | — | unavailable |
| POST | `v1/devices/{deviceId}/base/pairfirstfoundbase` | Pair first found base | — | — | unavailable |

---

## Household

**Source:** `jadx-out/sources/Nb/InterfaceC6136a.java`

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v1/household/users/{userId}/summary` | Get household summary | — | unavailable |
| GET | `v1/household/users/{userId}/invitations` | Get invitations | — | unavailable |
| POST | `v1/household/households/{householdId}/users` | Invite user | `InviteUserToHouseholdRequest` | unavailable |
| POST | `v1/household/households/{householdId}/users/{userId}` | Respond to invitation | `InvitationResponseRequest` | unavailable |
| DELETE | `v1/household/households/{householdId}/users/{guestUserId}` | Remove guest | — | unavailable |
| POST | `v1/household/households/{householdId}/devices` | Add device to household | `AddHouseholdDeviceRequest` | unavailable |
| DELETE | `v1/household/devices/{deviceId}` | Remove device | — | unavailable |
| PUT | `v1/household/devices/{deviceId}` | Update device info | `HouseholdUpdateDeviceInfoRequest` | unavailable |
| POST | `v1/household/households/{householdId}/devices/{deviceId}/guests` | Add guest to device | `AddGuestRequest` | unavailable |
| DELETE | `v1/household/households/{householdId}/sets/{setId}` | Delete device set | — | unavailable |
| PUT | `v1/household/households/{householdId}/sets/{setId}` | Update device set | `HouseholdUpdateDeviceSetRequest` | unavailable |
| GET | `v1/household/users/{userId}/schedule/{setId}` | Get schedule for set | — | unavailable |
| DELETE | `v1/household/users/{userId}/schedule/{setId}` | Delete schedule | — | unavailable |
| POST | `v1/household/users/{userId}/schedule` | Create/update schedule | `ReturnDateRequest` | unavailable |
| GET | `v1/household/users/{userId}/current-set` | Get current device set | — | unavailable |
| PUT | `v1/household/users/{userId}/current-set` | Set current device set | `SetCurrentDeviceSetRequest` | unavailable |
| DELETE | `v1/household/users/{userId}/current-set` | Clear current device set | — | unavailable |

---

## Travel / Jetlag

**Source:** `jadx-out/sources/Db/a.java`, `Db/b.java`

### Trips

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v1/users/{userId}/travel/trips` | List all trips | — | unavailable |
| POST | `v1/users/{userId}/travel/trips` | Create trip | `CreateTripRequest` | unavailable |
| GET | `v1/users/{userId}/travel/trips/{tripId}` | Get trip details | — | unavailable |
| PUT | `v1/users/{userId}/travel/trips/{tripId}` | Update trip | `UpdateTripRequest` | unavailable |
| DELETE | `v1/users/{userId}/travel/trips/{tripId}` | Delete trip | — | unavailable |

### Plans

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| POST | `v1/users/{userId}/travel/trips/{tripId}/plans` | Create plan | `CreateJetLagPlanRequest` | unavailable |
| GET | `v1/users/{userId}/travel/trips/{tripId}/plans` | Get plans for trip | — | unavailable |
| PATCH | `v1/users/{userId}/travel/plans/{planId}/tasks` | Update plan tasks | `JetLagBulkTaskUpdateRequest` | unavailable |

### Utilities

| Method | Path | Purpose | Query Params | Status |
|--------|------|---------|--------------|--------|
| GET | `v1/travel/airport-search` | Search airports | `query`, `maxResults` | unavailable |
| GET | `v1/travel/flight-status` | Get flight status | `flightNumber`, `date` | unavailable |

---

## Insights

### Daily Insights

**Source:** `jadx-out/sources/vb/InterfaceC7510a.java`

| Method | Path | Purpose | Query Params | Status |
|--------|------|---------|--------------|--------|
| GET | `v1/users/{userId}/insights` | Get daily insights | `date` (YYYY-MM-DD) | unavailable |

### AI (LLM) Insights

**Source:** `jadx-out/sources/Z8/a.java`

| Method | Path | Purpose | Query Params | Body | Status |
|--------|------|---------|--------------|------|--------|
| GET | `v1/users/{userId}/llm-insights` | Get AI insights | `from`, `to` (date) | — | unavailable |
| POST | `v1/users/{userId}/llm-insights/batch` | Get batch insights | — | `AiInsightsBatchRequest` | unavailable |
| GET | `v1/users/{userId}/llm-insights/settings` | Get AI insight settings | — | — | unavailable |
| PUT | `v1/users/{userId}/llm-insights/settings` | Update AI insight settings | — | `AiInsightsSettingsUpdateRequest` | unavailable |
| POST | `v1/users/{userId}/llm-insights/{insightId}/feedback` | Post insight feedback | — | `AiInsightFeedbackRequest` | unavailable |

---

## Health Integrations

**Source:** `jadx-out/sources/Va/a.java`

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| POST | `v1/users/{userId}/health-integrations/sources/{sourceId}` | Upload health data | `HealthIntegrationUploadRequest` | unavailable |
| GET | `v1/users/{userId}/health-integrations/sources/{sourceId}/checkpoints` | Get health checkpoints | — | unavailable |
| GET | `v1/users/{userId}/health-integrations/metadata` | Get health metadata | — | unavailable |

---

## Notifications

**Source:** `jadx-out/sources/Zb/a.java`

| Method | Path | Purpose | Query Params | Body | Status |
|--------|------|---------|--------------|------|--------|
| GET | `v1/users/{userId}/notifications` | Get active notifications | `active=true` | — | unavailable |
| POST | `v1/push_event/acknowledge` | Acknowledge notification | — | `NotificationAckRequest` | unavailable |

---

## Challenges

**Source:** `jadx-out/sources/Ca/InterfaceC3050a.java`

| Method | Path | Purpose | Query Params | Status |
|--------|------|---------|--------------|--------|
| GET | `v1/users/{userId}/challenges` | Get user challenges | `state` (optional) | unavailable |

---

## Bedtime Schedule

**Source:** `jadx-out/sources/T9/a.java`

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v1/users/{userId}/bedtime/recommendation` | Get bedtime recommendation | — | unavailable |
| PUT | `v1/users/{userId}/bedtime` | Update bedtime schedule | `BedtimeScheduleSettingsRequest` | unavailable |

---

## Presence / Truth Tags

**Source:** `jadx-out/sources/Qc/a.java`

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v1/users/{userId}/truth-tags` | Get all presence tags | — | unavailable |
| POST | `v1/users/{userId}/truth-tags` | Create tag | `PresenceTagsRequest` | unavailable |
| PUT | `v1/users/{userId}/truth-tags/{tagId}` | Update tag | `PresenceTagsRequest` | unavailable |
| DELETE | `v1/users/{userId}/truth-tags/{tagId}` | Delete tag | — | unavailable |

---

## Onboarding

**Source:** `jadx-out/sources/Gc/InterfaceC4670a.java`

| Method | Path | Purpose | Body | Status |
|--------|------|---------|------|--------|
| GET | `v1/users/{userId}/app-state/onboard` | Get onboarding state | — | unavailable |
| PUT | `v1/users/{userId}/app-state/onboard` | Update onboarding state | `RemoteOnboardingStateNetworkModel` | unavailable |

---

## Validation Status Summary

| Domain | Available | Unavailable |
|--------|-----------|-------------|
| Auth | 2 | 2 |
| Devices | 6 | 11 |
| Users & Profile | 1 | 0 |
| Temperature | 0 | 13 |
| Alarms | 0 | 10 |
| Sleep & Metrics | 4 | 3 |
| Autopilot | 0 | 6 |
| Adjustable Base | 0 | 8 |
| Household | 0 | 17 |
| Travel / Jetlag | 0 | 10 |
| Insights (daily + AI) | 0 | 6 |
| Health Integrations | 0 | 3 |
| Notifications | 0 | 2 |
| Challenges | 0 | 1 |
| Bedtime | 0 | 2 |
| Presence Tags | 0 | 4 |
| Onboarding | 0 | 2 |
| **Total** | **13** | **100** |

## Validation Order (per INSTRUCTIONS.md)

1. `POST v1/tokens` — Auth
2. `GET v1/users/me` — Bootstrap
3. `GET v1/devices/{deviceId}` — Device read
4. `GET v1/users/{userId}/temperature/all` — Temperature read
5. `GET v1/users/{userId}/trends` — Sleep trends
6. `GET v1/household/users/{userId}/summary` — Household read
7. `GET v1/users/{userId}/travel/trips` — Travel read
8. Write endpoints last, one at a time

---

*Originally generated from static APK analysis; statuses above were live-validated with `curl` against the real API on 2026-03-14.*

## Why So Many Endpoints Are Unavailable

The dominant failure mode during live validation was an HTML response of the form `Cannot METHOD /v1/...`. That strongly suggests route mismatch rather than auth failure or bad request payloads. The most likely causes are:

1. The APK-derived paths are stale and the mobile app has moved those features to different routes since the version that was decompiled.
2. Some features may now live on different hosts or subdomains than `client-api.8slp.net`.
3. Some endpoints may require a different API version prefix such as `/v2` or a non-REST transport.
4. Some features may be behind product, region, firmware, account-tier, or device-capability gates, with the app calling different routes conditionally.
5. A subset of write routes may still exist but require different request shapes, though that would not explain the large number of `Cannot METHOD /path` responses by itself.

## How To Identify The Correct Endpoints

1. Re-run the APK analysis on the latest Android app build rather than the older December 2025 assumptions.
2. Trace the exact request builder or Retrofit interface for each unavailable feature and confirm host, method, version, path, query params, and body model.
3. Inspect interceptors and base URL providers for host switching, feature flags, and environment selection.
4. Search for alternate strings around each failed domain, especially `v2`, `v3`, `graphql`, `gateway`, `journey`, `bedtime`, `jetlag`, `autopilot`, and `household`.
5. Capture real app traffic from the current mobile app with a trusted proxy setup and compare actual requests against the audited paths.
6. Compare account bootstrap payloads from `GET /v1/users/me` to determine which features are actually enabled on this account and device.
7. For routes that partly overlap with working behavior, start from known-good calls like `GET /v1/users/{userId}/trends` and inspect nearby code paths in the app for related endpoints.

## How To Rectify And Fix The Endpoints

1. Treat every `unavailable` route in this file as stale until re-proven by current runtime evidence.
2. Update the Go client and this audit only after confirming an exact live method, host, and path from either current app traffic or current APK code.
3. Split future validation into three classes: route exists, route exists but request schema is wrong, and route is feature-gated for this account.
4. For mutating endpoints, validate with intentionally invalid bodies first; if the API returns structured JSON validation errors instead of `Cannot METHOD /path`, the route likely exists.
5. Add host/version notes per endpoint once re-discovered so the repo does not implicitly assume every feature lives under `https://client-api.8slp.net/v1`.
6. Prefer replacing stale routes incrementally by domain: alarms, temperature modes, household, travel, then autopilot and insights.
7. Keep a small set of known-good discovery calls at the top of the workflow: `POST /v1/tokens`, `GET /v1/users/me`, `GET /v1/devices/{deviceId}`, and `GET /v1/users/{userId}/trends`.
