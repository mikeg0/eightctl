# Eight Sleep APK Reverse Engineering Instructions

## Purpose

This file is for an LLM or engineer reverse engineering the Eight Sleep Android APK and validating undocumented API endpoints for use in this repository.

The objective is to:

1. Extract and verify endpoint definitions from decompiled APK output.
2. Compare those endpoints to the current Go client implementation.
3. Test endpoints safely on the user's own account and device.
4. Update this repository only after an endpoint has been validated.

Do not guess endpoint details when the APK or runtime evidence is ambiguous.

## Scope

Focus on these domains:

1. Auth
2. User and device discovery
3. Temperature and power
4. Sleep and metrics
5. Household
6. Audio
7. Base
8. Travel

## Inputs

Assume these inputs may exist locally:

1. `base.apk`
2. split APKs such as `split_config.arm64_v8a.apk`, `split_config.en.apk`, `split_config.xxhdpi.apk`
3. Decompiled output from `jadx`
4. Decoded resources from `apktool`
5. The current repository source code

## Working Directories

Use a workspace like this:

```text
~/apk-work/
  base.apk
  jadx-out/
  apktool-out/
```

## Required Tools

Expected tools:

1. `jadx`
2. `apktool`
3. `rg`
4. `adb` optionally

## High-Level Workflow

Follow this sequence:

1. Decompile `base.apk` with `jadx`.
2. Decode `base.apk` with `apktool`.
3. Determine whether the app logic is Java/Kotlin, React Native, Flutter, or hybrid.
4. Find auth, base URLs, and common headers first.
5. Enumerate endpoint definitions and group them by domain.
6. Trace request and response models.
7. Build an endpoint inventory.
8. Compare that inventory to the Go client in this repo.
9. Validate endpoints using safe read-only tests first.
10. Only after validation, patch repository code.

## Decompile

Run:

```bash
jadx -d jadx-out base.apk
apktool d base.apk -o apktool-out
```

If split APKs exist, do not start by decompiling them unless code or strings appear missing from `base.apk`.

## Determine App Architecture

Before tracing endpoints, determine where the business logic lives.

Run:

```bash
rg -n "retrofit2\.http|OkHttpClient|Interceptor|index.android.bundle|hermes|flutter_assets|libapp\.so" jadx-out apktool-out
```

Interpretation:

1. If Retrofit and OkHttp classes are present, network logic is likely in Java or Kotlin.
2. If `index.android.bundle` or Hermes artifacts dominate, inspect React Native assets.
3. If `flutter_assets` or `libapp.so` dominate, inspect Flutter artifacts and string tables.

If the main logic is not in Java or Kotlin, do not assume missing endpoints do not exist. Switch analysis to the actual runtime bundle.

## Search Strategy

Start with these searches:

```bash
rg -n "auth-api|client-api|8slp|client_id|client_secret|grant_type|Authorization|Bearer|User-Agent|Content-Type" jadx-out apktool-out
rg -n "@(GET|POST|PUT|DELETE|PATCH)|retrofit2\.http|Request\.Builder|newCall|/v1/|/users/|/devices/" jadx-out
rg -n "travel|trip|plan|airport|flight|metrics|summary|aggregate|trends|insights" jadx-out apktool-out
```

The first task is to find:

1. Auth host
2. API host
3. Auth payload fields
4. Default client credentials if present
5. Common headers
6. API version markers such as `/v1`

## How To Trace Endpoints

When a path or host string is found:

1. Use cross references in `jadx-gui`.
2. Find the method that uses the string.
3. Record the HTTP method.
4. Record the host and path.
5. Record query parameters.
6. Record request body model.
7. Record response model.
8. Record any required auth token or derived ID.

If annotations such as `@GET` or `@POST` are present, treat those as the primary source of truth.

If requests are built manually, trace:

1. URL construction
2. Headers
3. Body serialization
4. Interceptors
5. Retry logic

## Endpoint Inventory Format

Maintain an inventory table with these columns:

1. Domain
2. Method
3. Host
4. Path template
5. Query params
6. Body fields
7. Auth type
8. Required IDs
9. Response model
10. Source file and class
11. Confidence
12. Validation status

Confidence values:

1. `high`: explicit route annotation or direct request builder found
2. `medium`: inferred from model and nearby code
3. `low`: string-only evidence or unclear host

Validation status values:

1. `unverified`
2. `verified-read`
3. `verified-write`
4. `stale`
5. `failed`

## Required IDs

Track how each endpoint gets its identifiers.

Common IDs to resolve:

1. `userId`
2. `deviceId`
3. `householdId`
4. travel-specific IDs such as trip or plan IDs

Do not hardcode guessed IDs. Trace where the app obtains them, usually from `/users/me` or similar bootstrap endpoints.

## Compare Against This Repository

Compare the endpoint inventory to these areas in the repo:

1. `internal/client/eightsleep.go`
2. `internal/client/metrics.go`
3. `internal/client/travel.go`
4. other files under `internal/client/`

For each route in the Go client, classify it:

1. `matches APK`
2. `partial mismatch`
3. `missing in APK evidence`
4. `new in APK`

Do not patch code until a route is at least `high` or `medium` confidence and there is a plan to validate it.

## Runtime Validation Order

Validate in this order:

1. Auth
2. `/users/me` or equivalent bootstrap route
3. read-only user and device routes
4. read-only temperature and status routes
5. sleep and trends routes
6. household routes
7. travel read routes
8. reversible write routes such as power and temperature
9. other write routes last

Never start with write endpoints.

## Safety Rules For Testing

Use only the user's own account and device.

Avoid:

1. high-frequency loops
2. brute-force endpoint guessing
3. repeated login attempts
4. write tests that change alarms, schedules, or travel state unless the route is understood and reversible

Prefer:

1. reusing cached tokens
2. spacing requests
3. testing one route at a time
4. capturing exact failure details

## How To Test Endpoints

Use the repository's client as the test harness when possible because it already includes auth and app-like headers.

Before broad testing, add temporary request tracing in the shared HTTP path:

1. method
2. full URL
3. query parameters
4. redacted headers
5. redacted body
6. response status
7. response body on non-2xx

Redact:

1. bearer tokens
2. passwords
3. client secrets

If a new endpoint is discovered from the APK, add a temporary client method and test it through the same shared HTTP path instead of creating separate ad hoc scripts.

## Error Interpretation

Use these heuristics:

1. `401`: token or auth style is wrong, or required headers are missing
2. `403`: account lacks access or feature is gated
3. `404` HTML such as `Cannot GET /v1/...`: wrong host, wrong path, or wrong API version
4. `404` JSON: route exists but resource ID is wrong
5. `405`: method is wrong
6. `400`: body or query params are wrong
7. `415`: content type is wrong
8. `429`: too many requests, back off and stop retry loops

## Travel-Specific Instructions

For travel, search for:

```bash
rg -n "travel|trip|plan|airport|flight|itinerary|destination|timezone" jadx-out apktool-out
```

For each travel hit:

1. verify whether it is a real network route or only UI text
2. identify the surrounding request builder or Retrofit interface
3. record host, method, path, params, and body
4. determine how travel IDs are discovered
5. validate read endpoints before write endpoints

Do not assume the travel routes currently in `internal/client/travel.go` are still valid.

## Metrics-Specific Instructions

For metrics, search for:

```bash
rg -n "metrics|summary|aggregate|trends|intervals|insights|sleepQualityScore" jadx-out apktool-out
```

For each metrics hit:

1. confirm the exact host and path
2. confirm query parameters such as `from`, `to`, `model-version`, or any feature flags
3. compare against `internal/client/metrics.go`
4. prefer routes clearly used by the app over guessed routes

If a route returns HTML `Cannot GET`, treat it as stale until proven otherwise.

## When Static Analysis Is Not Enough

If the APK reveals a domain and models but not an unambiguous route:

1. inspect interceptors and base URL providers
2. inspect assets and bundles
3. inspect string tables
4. inspect manifest and network config
5. use runtime observation only as a tie-breaker after static analysis

Do not jump directly to runtime testing if the static evidence is weak and the route is write-capable.

## Deliverables

A complete analysis pass should produce:

1. an endpoint inventory
2. a list of mismatches between APK and Go client
3. a list of validated routes
4. a list of stale or removed routes
5. a minimal patch plan for this repo

## Patch Rules

When updating this repository:

1. change one domain at a time
2. keep auth and shared request logic centralized
3. prefer small patches
4. add or update tests where practical
5. do not remove old routes without evidence they are stale
6. mention whether a route was confirmed from static analysis, runtime validation, or both

## Stop Conditions

Stop and report instead of guessing when:

1. the route is only implied by UI strings
2. the host is unknown
3. the method is unknown
4. the required ID source is unknown
5. the endpoint is write-capable and not safely reversible

## Minimal Command Set

Use this command set first:

```bash
jadx -d jadx-out base.apk
apktool d base.apk -o apktool-out
rg -n "auth-api|client-api|8slp|client_id|client_secret|grant_type|Authorization|Bearer" jadx-out apktool-out
rg -n "@(GET|POST|PUT|DELETE|PATCH)|retrofit2\.http|Request\.Builder|newCall|/v1/|/users/|/devices/" jadx-out
rg -n "travel|trip|plan|airport|flight|metrics|summary|aggregate|trends|insights" jadx-out apktool-out
```

## Final Rule

Prefer evidence over intuition. If the APK, the runtime behavior, and the current Go client disagree, trust the strongest direct evidence and document the mismatch explicitly.
