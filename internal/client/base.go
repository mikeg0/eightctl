package client

import (
	"context"
	"fmt"
	"net/http"
)

type BaseActions struct{ c *Client }

func (c *Client) Base() *BaseActions { return &BaseActions{c: c} }

// Info fetches adjustable base state from app-api.
func (b *BaseActions) Info(ctx context.Context) (any, error) {
	if err := b.c.requireUser(ctx); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/v1/users/%s/base", b.c.UserID)
	var res any
	err := b.c.doApp(ctx, http.MethodGet, path, nil, nil, &res)
	return res, err
}

// SetAngle sets the adjustable base head/foot angles via app-api.
func (b *BaseActions) SetAngle(ctx context.Context, head, foot int) error {
	if err := b.c.requireUser(ctx); err != nil {
		return err
	}
	path := fmt.Sprintf("/v1/users/%s/base/angle?ignoreDeviceErrors=false", b.c.UserID)
	body := map[string]any{"head": head, "foot": foot}
	return b.c.doApp(ctx, http.MethodPost, path, nil, body, nil)
}
