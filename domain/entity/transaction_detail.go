package entity

import "time"

type TransactionDetail struct {
	ID                int
	TransactionID     int
	InstallmentAmount float64
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

func (TransactionDetail) TableName() string {
	return "d_transaction_details"
}
