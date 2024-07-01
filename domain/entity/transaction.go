package entity

import "time"

type Transaction struct {
	ID                int
	ConsumerID        int
	ContractNumber    string
	OTR               float64
	AdminFee          float64
	InstallmentAmount float64
	InterestAmount    float64
	AssetName         string
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

func (Transaction) TableName() string {
	return "d_transactions"
}
