package repository

import (
	"context"
	"devices/models"
	"errors"
)

var (
	ErrDeviceNotFound = errors.New("device not found")
)

type Repository interface {
	AddDevice(ctx context.Context, device models.Device) (models.Device, error)
	GetDeviceById(ctx context.Context, id int32) (models.Device, error)
	UpdateDevice(ctx context.Context, device models.Device) error
}
