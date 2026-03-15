package client

import (
	"context"
	"fmt"
	"net/http"
)

type HouseholdActions struct{ c *Client }

func (c *Client) Household() *HouseholdActions { return &HouseholdActions{c: c} }

// Summary fetches household summary from app-api.
func (h *HouseholdActions) Summary(ctx context.Context) (any, error) {
	if err := h.c.requireUser(ctx); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/v1/household/users/%s/summary", h.c.UserID)
	var res any
	err := h.c.doApp(ctx, http.MethodGet, path, nil, nil, &res)
	return res, err
}
