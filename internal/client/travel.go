package client

import (
	"context"
	"fmt"
	"net/http"
)

type TravelActions struct{ c *Client }

func (c *Client) Travel() *TravelActions { return &TravelActions{c: c} }

// Trips lists travel trips from app-api.
func (t *TravelActions) Trips(ctx context.Context) (any, error) {
	if err := t.c.requireUser(ctx); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/v1/users/%s/travel/trips", t.c.UserID)
	var res any
	err := t.c.doApp(ctx, http.MethodGet, path, nil, nil, &res)
	return res, err
}

// CreateTrip creates a new trip on app-api.
func (t *TravelActions) CreateTrip(ctx context.Context, body map[string]any) error {
	if err := t.c.requireUser(ctx); err != nil {
		return err
	}
	path := fmt.Sprintf("/v1/users/%s/travel/trips", t.c.UserID)
	return t.c.doApp(ctx, http.MethodPost, path, nil, body, nil)
}

// DeleteTrip deletes a trip on app-api.
func (t *TravelActions) DeleteTrip(ctx context.Context, tripID string) error {
	if err := t.c.requireUser(ctx); err != nil {
		return err
	}
	path := fmt.Sprintf("/v1/users/%s/travel/trips/%s", t.c.UserID, tripID)
	return t.c.doApp(ctx, http.MethodDelete, path, nil, nil, nil)
}
