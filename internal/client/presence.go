package client

import (
	"context"
	"fmt"
	"net/http"
)

// GetPresence derives bed presence from the device data (leftPresenceStart/rightPresenceStart).
// The /presence endpoint is unavailable on both hosts; presence is inferred from device polling.
func (c *Client) GetPresence(ctx context.Context) (bool, error) {
	id, err := c.EnsureDeviceID(ctx)
	if err != nil {
		return false, err
	}
	path := fmt.Sprintf("/devices/%s", id)
	var res struct {
		Result struct {
			LeftPresenceStart  float64 `json:"leftPresenceStart"`
			LeftPresenceEnd    float64 `json:"leftPresenceEnd"`
			RightPresenceStart float64 `json:"rightPresenceStart"`
			RightPresenceEnd   float64 `json:"rightPresenceEnd"`
			LeftUserId         string  `json:"leftUserId"`
			RightUserId        string  `json:"rightUserId"`
		} `json:"result"`
	}
	if err := c.do(ctx, http.MethodGet, path, nil, nil, &res); err != nil {
		return false, err
	}
	// Determine which side the user is on and check if presence is active.
	if c.UserID == res.Result.LeftUserId {
		return res.Result.LeftPresenceEnd == 0 && res.Result.LeftPresenceStart > 0, nil
	}
	if c.UserID == res.Result.RightUserId {
		return res.Result.RightPresenceEnd == 0 && res.Result.RightPresenceStart > 0, nil
	}
	return false, fmt.Errorf("user %s not assigned to a device side", c.UserID)
}
