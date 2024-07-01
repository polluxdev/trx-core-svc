package serializer

import (
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type ConsumerDTO struct {
	ID           int     `json:"id"`
	NIK          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
	IdCardPhoto  string  `json:"id_card_photo"`
	SelfiePhoto  string  `json:"selfie_photo"`
	CreatedAt    string  `json:"created_at"`
}

func SerializeConsumer(data *entity.Consumer) *ConsumerDTO {
	result := &ConsumerDTO{
		ID:           data.ID,
		NIK:          data.NIK,
		FullName:     data.FullName,
		LegalName:    data.LegalName,
		PlaceOfBirth: data.PlaceOfBirth,
		DateOfBirth:  helper.TimeToString(&data.DateOfBirth, global.DATE_FORMAT),
		Salary:       data.Salary,
		IdCardPhoto:  data.IdCardPhoto,
		SelfiePhoto:  data.SelfiePhoto,
		CreatedAt:    helper.TimeToString(data.CreatedAt, global.DATE_TIME_FORMAT),
	}

	return result
}

func SerializeConsumers(data []*entity.Consumer) []*ConsumerDTO {
	result := make([]*ConsumerDTO, 0)
	for _, item := range data {
		result = append(result, SerializeConsumer(item))
	}

	return result
}
