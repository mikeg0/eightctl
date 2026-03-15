package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AlarmActions groups alarm endpoints on app-api.
type AlarmActions struct {
	c *Client
}

// Alarms helper accessor.
func (c *Client) Alarms() *AlarmActions { return &AlarmActions{c: c} }

func (a *AlarmActions) Snooze(ctx context.Context, alarmID string) error {
	if err := a.c.requireUser(ctx); err != nil {
		return err
	}
	path := fmt.Sprintf("/v1/users/%s/alarms/%s/snooze", a.c.UserID, alarmID)
	return a.c.doApp(ctx, http.MethodPut, path, nil, nil, nil)
}

// Dismiss disables an alarm by setting enabled=false via the routines API.
func (a *AlarmActions) Dismiss(ctx context.Context, alarmID string) error {
	return a.setAlarmEnabled(ctx, alarmID, false)
}

// DismissAll disables all alarms.
func (a *AlarmActions) DismissAll(ctx context.Context) error {
	routines, err := a.getRoutines(ctx)
	if err != nil {
		return err
	}
	alarms, err := a.extractOneOffAlarms(routines)
	if err != nil {
		return err
	}
	for _, alarm := range alarms {
		alarmMap, ok := alarm.(map[string]any)
		if !ok {
			continue
		}
		if enabled, _ := alarmMap["enabled"].(bool); enabled {
			alarmMap["enabled"] = false
		}
	}
	return a.putRoutines(ctx, routines)
}

func (a *AlarmActions) VibrationTest(ctx context.Context) error {
	if err := a.c.requireUser(ctx); err != nil {
		return err
	}
	path := fmt.Sprintf("/v1/users/%s/vibration-test", a.c.UserID)
	return a.c.doApp(ctx, http.MethodPut, path, nil, nil, nil)
}

// setAlarmEnabled toggles enabled on a single alarm by ID via the routines API.
func (a *AlarmActions) setAlarmEnabled(ctx context.Context, alarmID string, enabled bool) error {
	routines, err := a.getRoutines(ctx)
	if err != nil {
		return err
	}
	alarms, err := a.extractOneOffAlarms(routines)
	if err != nil {
		return err
	}
	found := false
	for _, alarm := range alarms {
		alarmMap, ok := alarm.(map[string]any)
		if !ok {
			continue
		}
		if alarmMap["alarmId"] == alarmID {
			alarmMap["enabled"] = enabled
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("alarm %s not found", alarmID)
	}
	return a.putRoutines(ctx, routines)
}

func (a *AlarmActions) getRoutines(ctx context.Context) (map[string]any, error) {
	if err := a.c.requireUser(ctx); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/v2/users/%s/routines", a.c.UserID)
	var res map[string]any
	if err := a.c.doApp(ctx, http.MethodGet, path, nil, nil, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AlarmActions) putRoutines(ctx context.Context, routines map[string]any) error {
	path := fmt.Sprintf("/v2/users/%s/routines", a.c.UserID)
	// Extract just the settings to PUT back.
	settings, ok := routines["settings"].(map[string]any)
	if !ok {
		return fmt.Errorf("missing settings in routines")
	}
	return a.c.doApp(ctx, http.MethodPut, path, nil, settings, nil)
}

func (a *AlarmActions) extractOneOffAlarms(routines map[string]any) ([]any, error) {
	settings, ok := routines["settings"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("missing settings in routines")
	}
	alarms, ok := settings["oneOffAlarms"].([]any)
	if !ok {
		// Try to decode from json.RawMessage or similar
		if raw, ok := settings["oneOffAlarms"]; ok {
			b, err := json.Marshal(raw)
			if err != nil {
				return nil, fmt.Errorf("cannot read oneOffAlarms")
			}
			var arr []any
			if err := json.Unmarshal(b, &arr); err != nil {
				return nil, fmt.Errorf("cannot parse oneOffAlarms")
			}
			return arr, nil
		}
		return nil, fmt.Errorf("no oneOffAlarms found")
	}
	return alarms, nil
}
