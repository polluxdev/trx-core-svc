package serializer

import (
	"testing"
	"time"

	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"github.com/polluxdev/trx-core-svc/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestSerializeConsumer(t *testing.T) {
	dateOfBirth := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	createdAt := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	consumer := &entity.Consumer{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "Windah Basudara",
		LegalName:    "Brando Windah Franco",
		PlaceOfBirth: "Jakarta",
		DateOfBirth:  dateOfBirth,
		Salary:       50000,
		IdCardPhoto:  "id_card_photo_url",
		SelfiePhoto:  "selfie_photo_url",
		CreatedAt:    &createdAt,
	}

	expected := &ConsumerDTO{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "Windah Basudara",
		LegalName:    "Brando Windah Franco",
		PlaceOfBirth: "Jakarta",
		DateOfBirth:  helper.TimeToString(&dateOfBirth, global.DATE_FORMAT),
		Salary:       50000,
		IdCardPhoto:  "id_card_photo_url",
		SelfiePhoto:  "selfie_photo_url",
		CreatedAt:    helper.TimeToString(&createdAt, global.DATE_TIME_FORMAT),
	}

	result := SerializeConsumer(consumer)
	assert.Equal(t, expected, result)
}

func TestSerializeConsumers(t *testing.T) {
	dateOfBirth1 := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	createdAt1 := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	dateOfBirth2 := time.Date(1985, time.March, 15, 0, 0, 0, 0, time.UTC)
	createdAt2 := time.Date(2022, time.December, 31, 0, 0, 0, 0, time.UTC)

	consumers := []*entity.Consumer{
		{
			ID:           1,
			NIK:          "1234567890",
			FullName:     "Windah Basudara",
			LegalName:    "Brando Windah Franco",
			PlaceOfBirth: "Jakarta",
			DateOfBirth:  dateOfBirth1,
			Salary:       50000,
			IdCardPhoto:  "id_card_photo_url",
			SelfiePhoto:  "selfie_photo_url",
			CreatedAt:    &createdAt1,
		},
		{
			ID:           2,
			NIK:          "0987654321",
			FullName:     "Ilham Kurniawan",
			LegalName:    "Ilham Kurniawan Kurniadi",
			PlaceOfBirth: "Bekasi",
			DateOfBirth:  dateOfBirth2,
			Salary:       60000,
			IdCardPhoto:  "id_card_photo_url_2",
			SelfiePhoto:  "selfie_photo_url_2",
			CreatedAt:    &createdAt2,
		},
	}

	expected := []*ConsumerDTO{
		{
			ID:           1,
			NIK:          "1234567890",
			FullName:     "Windah Basudara",
			LegalName:    "Brando Windah Franco",
			PlaceOfBirth: "Jakarta",
			DateOfBirth:  helper.TimeToString(&dateOfBirth1, global.DATE_FORMAT),
			Salary:       50000,
			IdCardPhoto:  "id_card_photo_url",
			SelfiePhoto:  "selfie_photo_url",
			CreatedAt:    helper.TimeToString(&createdAt1, global.DATE_TIME_FORMAT),
		},
		{
			ID:           2,
			NIK:          "0987654321",
			FullName:     "Ilham Kurniawan",
			LegalName:    "Ilham Kurniawan Kurniadi",
			PlaceOfBirth: "Bekasi",
			DateOfBirth:  helper.TimeToString(&dateOfBirth2, global.DATE_FORMAT),
			Salary:       60000,
			IdCardPhoto:  "id_card_photo_url_2",
			SelfiePhoto:  "selfie_photo_url_2",
			CreatedAt:    helper.TimeToString(&createdAt2, global.DATE_TIME_FORMAT),
		},
	}

	result := SerializeConsumers(consumers)
	assert.Equal(t, expected, result)
}
