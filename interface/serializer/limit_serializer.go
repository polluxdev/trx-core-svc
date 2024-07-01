package serializer

import (
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type LimitDTO struct {
	ID           int     `json:"id"`
	Duration     int     `json:"duration"`
	DurationType string  `json:"duration_type"`
	LimitAmount  float64 `json:"limit_amount"`
	CreatedAt    string  `json:"created_at"`
}

func SerializeLimit(data *entity.Limit) *LimitDTO {
	result := &LimitDTO{
		ID:           data.ID,
		Duration:     data.Duration,
		DurationType: data.DurationType,
		LimitAmount:  data.LimitAmount,
		CreatedAt:    *data.CreatedAt,
	}

	return result
}

func SerializeLimits(data []*entity.Limit) []*LimitDTO {
	result := make([]*LimitDTO, 0)
	for _, item := range data {
		result = append(result, SerializeLimit(item))
	}

	return result
}
