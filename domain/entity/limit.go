package entity

import "time"

type Limit struct {
	ID          int
	Duration    int
	LimitType   string
	LimitAmount float64
	Interest    float64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (Limit) TableName() string {
	return "md_limits"
}
