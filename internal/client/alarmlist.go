package client

import (
	"context"
	"fmt"
	"net/http"
)

// Alarm represents alarm payload.
type Alarm struct {
	ID         string  `json:"id"`
	Enabled    bool    `json:"enabled"`
	Time       string  `json:"time"`
	DaysOfWeek []int   `json:"daysOfWeek"`
	Vibration  bool    `json:"vibration"`
	Sound      *string `json:"sound,omitempty"`
}

// ListAlarms lists routines/alarms from app-api v2.
func (c *Client) ListAlarms(ctx context.Context) (any, error) {
	if err := c.requireUser(ctx); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/v2/users/%s/routines", c.UserID)
	var res any
	if err := c.doApp(ctx, http.MethodGet, path, nil, nil, &res); err != nil {
		return nil, err
	}
	return res, nil
}
