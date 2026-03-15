package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type MetricsActions struct{ c *Client }

func (c *Client) Metrics() *MetricsActions { return &MetricsActions{c: c} }

// Trends fetches daily sleep trends from client-api.
func (m *MetricsActions) Trends(ctx context.Context, from, to, tz string, out any) error {
	if err := m.c.requireUser(ctx); err != nil {
		return err
	}
	q := url.Values{}
	q.Set("from", from)
	q.Set("to", to)
	if tz != "" {
		q.Set("tz", tz)
	}
	q.Set("include-main", "false")
	q.Set("include-all-sessions", "true")
	q.Set("model-version", "v2")
	path := fmt.Sprintf("/users/%s/trends", m.c.UserID)
	return m.c.do(ctx, http.MethodGet, path, q, nil, out)
}

// Intervals fetches recent sleep intervals from client-api (no session ID required).
// Returns up to 10 intervals with full timeseries data; use cursor for pagination.
func (m *MetricsActions) Intervals(ctx context.Context, cursor string, out any) error {
	if err := m.c.requireUser(ctx); err != nil {
		return err
	}
	path := fmt.Sprintf("/users/%s/intervals", m.c.UserID)
	q := url.Values{}
	if cursor != "" {
		q.Set("next", cursor)
	}
	return m.c.do(ctx, http.MethodGet, path, q, nil, out)
}

// Summary fetches metrics summary from app-api.
func (m *MetricsActions) Summary(ctx context.Context, out any) error {
	if err := m.c.requireUser(ctx); err != nil {
		return err
	}
	q := url.Values{}
	q.Set("metrics", "all")
	path := fmt.Sprintf("/v1/users/%s/metrics/summary", m.c.UserID)
	return m.c.doApp(ctx, http.MethodGet, path, q, nil, out)
}

// Aggregate fetches aggregated metrics from app-api.
func (m *MetricsActions) Aggregate(ctx context.Context, out any) error {
	if err := m.c.requireUser(ctx); err != nil {
		return err
	}
	q := url.Values{}
	q.Set("metrics", "all")
	q.Set("v2", "true")
	path := fmt.Sprintf("/v1/users/%s/metrics/aggregate", m.c.UserID)
	return m.c.doApp(ctx, http.MethodGet, path, q, nil, out)
}

// Insights fetches daily insights from app-api.
func (m *MetricsActions) Insights(ctx context.Context, date string, out any) error {
	if err := m.c.requireUser(ctx); err != nil {
		return err
	}
	q := url.Values{}
	if date != "" {
		q.Set("date", date)
	}
	path := fmt.Sprintf("/v1/users/%s/insights", m.c.UserID)
	return m.c.doApp(ctx, http.MethodGet, path, q, nil, out)
}

// LLMInsights fetches AI-generated sleep insights from app-api.
func (m *MetricsActions) LLMInsights(ctx context.Context, from, to string, out any) error {
	if err := m.c.requireUser(ctx); err != nil {
		return err
	}
	q := url.Values{}
	if from != "" {
		q.Set("from", from)
	}
	if to != "" {
		q.Set("to", to)
	}
	path := fmt.Sprintf("/v1/users/%s/llm-insights", m.c.UserID)
	return m.c.doApp(ctx, http.MethodGet, path, q, nil, out)
}
