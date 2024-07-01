package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

func seedData(db *gorm.DB) {
	// Seed md_limits
	limits := []Limit{
		{ID: 1, Duration: 1, LimitType: "percentage", LimitAmount: 2, Interest: 0},
		{ID: 2, Duration: 2, LimitType: "percentage", LimitAmount: 4, Interest: 3.5},
		{ID: 3, Duration: 3, LimitType: "percentage", LimitAmount: 10, Interest: 5.8},
		{ID: 4, Duration: 4, LimitType: "percentage", LimitAmount: 14, Interest: 8.7},
	}
	for _, limit := range limits {
		db.FirstOrCreate(&limit)
	}

	// Seed md_consumers
	consumers := []Consumer{
		{ID: 1, NIK: "111", FullName: "Budi", LegalName: "Budi", PlaceOfBirth: "Jakarta", DateOfBirth: time.Date(2003, 10, 15, 0, 0, 0, 0, time.UTC), Salary: 5000000, IDCardPhoto: "111.jpg", SelfiePhoto: "111.jpg"},
		{ID: 2, NIK: "222", FullName: "Annisa", LegalName: "Annisa", PlaceOfBirth: "Bogor", DateOfBirth: time.Date(2004, 3, 21, 0, 0, 0, 0, time.UTC), Salary: 50000000, IDCardPhoto: "222.jpg", SelfiePhoto: "222.jpg"},
	}
	for _, consumer := range consumers {
		db.FirstOrCreate(&consumer)
	}

	// Seed md_consumer_limits
	consumerLimits := []ConsumerLimit{
		{ID: 1, ConsumerID: 1, LimitID: 1, Balance: 100000},
		{ID: 2, ConsumerID: 1, LimitID: 2, Balance: 200000},
		{ID: 3, ConsumerID: 1, LimitID: 3, Balance: 500000},
		{ID: 4, ConsumerID: 1, LimitID: 4, Balance: 700000},
		{ID: 5, ConsumerID: 2, LimitID: 1, Balance: 1000000},
		{ID: 6, ConsumerID: 2, LimitID: 2, Balance: 1200000},
		{ID: 7, ConsumerID: 2, LimitID: 3, Balance: 1500000},
		{ID: 8, ConsumerID: 2, LimitID: 4, Balance: 2000000},
	}
	for _, consumerLimit := range consumerLimits {
		db.FirstOrCreate(&consumerLimit)
	}
}
