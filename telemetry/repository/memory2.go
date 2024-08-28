package repository

import (
	"context"
	"sync"
	swagger "telemetry/go"
)

type MemoryRepository2 struct {
	mutex sync.RWMutex

	records map[string][]swagger.TelemetryRecord
}

func NewMemoryRepository2() *MemoryRepository2 {
	return &MemoryRepository2{
		records: make(map[string][]swagger.TelemetryRecord),
	}
}

func (r *MemoryRepository2) SetValue(ctx context.Context, value swagger.TelemetryRecord) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.records[value.Serial]; !ok {
		r.records[value.Serial] = make([]swagger.TelemetryRecord, 0)
	}

	r.records[value.Serial] = append(r.records[value.Serial], value)

	return nil
}

func (r *MemoryRepository2) GetValue(ctx context.Context, options RequestOptions) (swagger.InlineResponse200, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	perPage := max_records
	if options.PerPage < max_records {
		perPage = options.PerPage
	}
	page := options.Page
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * perPage
	result := make([]swagger.TelemetryRecord, 0)

	totalCount := len(r.records[options.Serial])

	if offset > totalCount {
		return swagger.InlineResponse200{
			Results: result,
			Total:   float64(totalCount),
			Page:    float64(page),
			PerPage: float64(perPage),
		}, nil
	}

	for i := offset; i < totalCount && i-offset < perPage; i++ {
		result = append(result, r.records[options.Serial][i])
	}

	return swagger.InlineResponse200{
		Results: result,
		Total:   float64(totalCount),
		Page:    float64(page),
		PerPage: float64(perPage),
	}, nil
}
