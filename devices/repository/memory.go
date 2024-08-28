package repository

import (
	"context"
	"devices/models"
	"sync"
)

type MemoryRepository struct {
	records map[int32]models.Device
	mutex   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		records: make(map[int32]models.Device),
	}
}

func (r *MemoryRepository) AddDevice(ctx context.Context, device models.Device) (models.Device, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id := int32(len(r.records) + 1)
	device.Id = int32(id)

	r.records[id] = device

	return device, nil
}

func (r *MemoryRepository) GetDeviceById(ctx context.Context, id int32) (models.Device, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	device, ok := r.records[id]
	if !ok {
		return models.Device{}, ErrDeviceNotFound
	}

	return device, nil
}

func (r *MemoryRepository) UpdateDevice(ctx context.Context, device models.Device) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, ok := r.records[device.Id]
	if !ok {
		return ErrDeviceNotFound
	}

	r.records[device.Id] = device

	return nil
}
