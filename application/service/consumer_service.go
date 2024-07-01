package service

import (
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/application/model"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"github.com/polluxdev/trx-core-svc/domain/entity"
	"github.com/polluxdev/trx-core-svc/infrastructure/repository"
)

type ConsumerService struct {
	consumerRepository repository.ConsumerRepositoryInterface
}

func NewConsumerService(consumerRepository repository.ConsumerRepositoryInterface) *ConsumerService {
	return &ConsumerService{consumerRepository: consumerRepository}
}

func (s *ConsumerService) CreateConsumer(request *model.NewConsumer) *entity.Consumer {
	data := &entity.Consumer{
		NIK:          request.NIK,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		PlaceOfBirth: request.PlaceOfBirth,
		DateOfBirth:  *helper.StringToTime(request.DateOfBirth, global.DATE_FORMAT),
		Salary:       request.Salary,
		IdCardPhoto:  request.IdCardPhoto,
		SelfiePhoto:  request.SelfiePhoto,
	}

	return s.consumerRepository.Store(data)
}

func (s *ConsumerService) GetConsumer(id int) *entity.Consumer {
	return s.consumerRepository.GetByColumn(&entity.Consumer{ID: id})
}
