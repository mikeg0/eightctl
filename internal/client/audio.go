package client

import (
	"context"
	"fmt"
	"net/http"
)

type AudioActions struct{ c *Client }

func (c *Client) Audio() *AudioActions { return &AudioActions{c: c} }

// Tracks lists audio tracks from app-api.
func (a *AudioActions) Tracks(ctx context.Context) ([]AudioTrack, error) {
	if err := a.c.requireUser(ctx); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/v1/users/%s/audio/tracks", a.c.UserID)
	var res struct {
		Tracks []AudioTrack `json:"tracks"`
	}
	if err := a.c.doApp(ctx, http.MethodGet, path, nil, nil, &res); err != nil {
		return nil, err
	}
	return res.Tracks, nil
}

// Categories lists audio categories from app-api.
func (a *AudioActions) Categories(ctx context.Context) (any, error) {
	path := "/v1/audio/categories"
	var res any
	err := a.c.doApp(ctx, http.MethodGet, path, nil, nil, &res)
	return res, err
}
