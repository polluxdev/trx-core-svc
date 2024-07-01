package entity

type ConsumerLimit struct {
	ID         int
	ConsumerID int
	LimitID    int
	CreatedAt  *string
}

func (ConsumerLimit) TableName() string {
	return "md_consumer_limits"
}
