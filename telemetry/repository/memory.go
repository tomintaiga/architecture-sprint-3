package repository

import (
	"context"
	"sync"
	swagger "telemetry/go"
	"time"
)

//-----------------------------------

func NewMemoryRepo() *MemoryRepository {
	return &MemoryRepository{
		records: make([]memRecord, 0),
	}
}

const (
	max_records = 10
)

type memRecord struct {
	created time.Time
	swagger.TelemetryRecord
}

type MemoryRepository struct {
	mutex sync.RWMutex

	records []memRecord
}

func (r *MemoryRepository) SetValue(ctx context.Context, value swagger.TelemetryRecord) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.records = append(r.records, memRecord{
		created:         time.Now(),
		TelemetryRecord: value,
	})

	return nil
}

func (r *MemoryRepository) GetValue(ctx context.Context, options RequestOptions) (swagger.InlineResponse200, error) {

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

	totalCount := len(r.records)

	if offset > totalCount {
		return swagger.InlineResponse200{
			Results: result,
			Total:   float64(totalCount),
			Page:    float64(page),
			PerPage: float64(perPage),
		}, nil
	}

	for i := offset; i < totalCount && i-offset < perPage; i++ {
		result = append(result, r.records[i].TelemetryRecord)
	}

	return swagger.InlineResponse200{
		Results: result,
		Total:   float64(totalCount),
		Page:    float64(page),
		PerPage: float64(perPage),
	}, nil
}
