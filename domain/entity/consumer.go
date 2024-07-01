package entity

import "time"

type Consumer struct {
	ID           int
	NIK          string
	FullName     string
	LegalName    string
	PlaceOfBirth string
	DateOfBirth  time.Time
	Salary       float64
	IdCardPhoto  string
	SelfiePhoto  string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (Consumer) TableName() string {
	return "md_consumers"
}
