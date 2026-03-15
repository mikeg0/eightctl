package client

import (
	"context"
	"fmt"
	"net/http"
)

type AutopilotActions struct{ c *Client }

func (c *Client) Autopilot() *AutopilotActions { return &AutopilotActions{c: c} }

// Details fetches autopilot configuration from app-api.
func (a *AutopilotActions) Details(ctx context.Context) (any, error) {
	if err := a.c.requireUser(ctx); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/v1/users/%s/autopilotDetails", a.c.UserID)
	var res any
	err := a.c.doApp(ctx, http.MethodGet, path, nil, nil, &res)
	return res, err
}
