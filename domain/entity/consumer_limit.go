package entity

import "time"

type ConsumerLimit struct {
	ID         int
	ConsumerID int
	LimitID    int
	Balance    float64
	CreatedAt  *time.Time

	Consumer Consumer `gorm:"foreignkey:ConsumerID;association_foreignkey:ID"`
	Limit    Limit    `gorm:"foreignkey:LimitID;association_foreignkey:ID"`
}

func (ConsumerLimit) TableName() string {
	return "md_consumer_limits"
}
