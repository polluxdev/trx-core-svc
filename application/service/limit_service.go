package service

import (
	"github.com/polluxdev/trx-core-svc/application/model"
	"github.com/polluxdev/trx-core-svc/domain/entity"
	"github.com/polluxdev/trx-core-svc/infrastructure/repository"
)

type LimitService struct {
	limitRepository repository.LimitRepositoryInterface
}

func NewLimitService(limitRepository repository.LimitRepositoryInterface) *LimitService {
	return &LimitService{limitRepository: limitRepository}
}

func (s *LimitService) GetLimitList(query *model.PaginationQuery) ([]*entity.Limit, *int64) {
	pagination := make(map[string]interface{})
	pagination["page"] = query.Page
	pagination["limit"] = query.Limit

	result, total := s.limitRepository.GetList(nil, nil, pagination)

	return result, total
}
