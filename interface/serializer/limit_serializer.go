package serializer

import (
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type LimitDTO struct {
	ID          string  `json:"id"`
	Duration    int     `json:"duration"`
	LimitType   string  `json:"limit_type"`
	LimitAmount float64 `json:"limit_amount"`
	CreatedAt   string  `json:"created_at"`
}

func SerializeLimit(data *entity.Limit) *LimitDTO {
	result := &LimitDTO{
		ID:          helper.Encrypt(data.ID),
		Duration:    data.Duration,
		LimitType:   data.LimitType,
		LimitAmount: data.LimitAmount,
		CreatedAt:   helper.TimeToString(data.CreatedAt, global.DATE_TIME_FORMAT),
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
