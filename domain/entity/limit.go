package entity

type Limit struct {
	ID           int
	Duration     int
	DurationType string
	LimitAmount  float64
	CreatedAt    *string
	UpdatedAt    *string
}

func (Limit) TableName() string {
	return "md_limits"
}
