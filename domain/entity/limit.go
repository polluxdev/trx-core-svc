package entity

type Limit struct {
	ID           int
	Duration     int
	DurationType string
	Limit        float64
	CreatedAt    *string
	UpdatedAt    *string
}

func (Limit) TableName() string {
	return "md_limits"
}
