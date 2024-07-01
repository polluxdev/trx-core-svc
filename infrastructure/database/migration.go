package database

import (
	"time"
)

type Consumer struct {
	ID           int       `gorm:"primaryKey"`
	NIK          string    `gorm:"size:255;not null"`
	FullName     string    `gorm:"size:255;not null"`
	LegalName    string    `gorm:"size:255;not null"`
	PlaceOfBirth string    `gorm:"size:255;not null"`
	DateOfBirth  time.Time `gorm:"not null"`
	Salary       float64   `gorm:"not null"`
	IDCardPhoto  string    `gorm:"size:255;not null"`
	SelfiePhoto  string    `gorm:"size:255;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (Consumer) TableName() string {
	return "md_consumers"
}

type Limit struct {
	ID          int       `gorm:"primaryKey"`
	Duration    int       `gorm:"not null"`
	LimitType   string    `gorm:"type:enum('fixed','percentage');not null"`
	LimitAmount float64   `gorm:"not null"`
	Interest    float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (Limit) TableName() string {
	return "md_limits"
}

type ConsumerLimit struct {
	ID         int       `gorm:"primaryKey"`
	ConsumerID int       `gorm:"not null"`
	LimitID    int       `gorm:"not null"`
	Balance    float64   `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	Consumer   Consumer  `gorm:"foreignKey:ConsumerID"`
	Limit      Limit     `gorm:"foreignKey:LimitID"`
}

func (ConsumerLimit) TableName() string {
	return "md_consumer_limits"
}

type Transaction struct {
	ID                int       `gorm:"primaryKey"`
	ConsumerID        int       `gorm:"not null"`
	ContractNumber    string    `gorm:"size:255;not null"`
	OTR               float64   `gorm:"not null"`
	AdminFee          float64   `gorm:"not null"`
	InstallmentAmount float64   `gorm:"not null"`
	InterestAmount    float64   `gorm:"not null"`
	AssetName         string    `gorm:"size:255;not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
	Consumer          Consumer  `gorm:"foreignKey:ConsumerID"`
}

func (Transaction) TableName() string {
	return "d_transactions"
}

type TransactionDetail struct {
	ID                int         `gorm:"primaryKey"`
	TransactionID     int         `gorm:"not null"`
	InstallmentAmount float64     `gorm:"not null"`
	CreatedAt         time.Time   `gorm:"autoCreateTime"`
	UpdatedAt         time.Time   `gorm:"autoUpdateTime"`
	Transaction       Transaction `gorm:"foreignKey:TransactionID"`
}

func (TransactionDetail) TableName() string {
	return "d_transaction_details"
}
