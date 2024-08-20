package repository

import (
	"context"
	"errors"
	swagger "telemetry/go"
)

type RequestOptions struct {
	Serial  string
	Page    int
	PerPage int
}

var (
	ErrDoesNotFound = errors.New("does not found")
)

type Repository interface {
	SetValue(ctx context.Context, value swagger.TelemetryRecord) error
	GetValue(ctx context.Context, options RequestOptions) (swagger.InlineResponse200, error)
}
