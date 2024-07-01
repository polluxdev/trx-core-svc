package service

import (
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/application/model"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"github.com/polluxdev/trx-core-svc/domain/entity"
	"github.com/polluxdev/trx-core-svc/infrastructure/repository"
)

type ConsumerService struct {
	ConsumerRepository repository.ConsumerRepositoryInterface
}

func NewConsumerService(ConsumerRepository repository.ConsumerRepositoryInterface) *ConsumerService {
	return &ConsumerService{ConsumerRepository: ConsumerRepository}
}

func (s *ConsumerService) CreateConsumer(request *model.NewConsumer) *entity.Consumer {
	Consumer := &entity.Consumer{
		IdCardNumber: request.IdCardNumber,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		PlaceOfBirth: request.PlaceOfBirth,
		DateOfBirth:  *helper.StringToTime(request.DateOfBirth, global.DATE_FORMAT),
		Salary:       request.Salary,
		IdCardPhoto:  request.IdCardPhoto,
		SelfiePhoto:  request.SelfiePhoto,
	}

	return s.ConsumerRepository.Store(Consumer)
}

func (s *ConsumerService) GetConsumer(id int) *entity.Consumer {
	return s.ConsumerRepository.GetByColumn(&entity.Consumer{ID: id})
}
