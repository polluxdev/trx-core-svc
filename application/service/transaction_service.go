package service

import (
	"fmt"

	"github.com/polluxdev/trx-core-svc/application/model"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"github.com/polluxdev/trx-core-svc/common/utils"
	"github.com/polluxdev/trx-core-svc/domain/entity"
	"github.com/polluxdev/trx-core-svc/infrastructure/repository"
)

type TransactionService struct {
	consumerLimitRepository     repository.ConsumerLimitRepositoryInterface
	transactionRepository       repository.TransactionRepositoryInterface
	transactionDetailRepository repository.TransactionDetailRepositoryInterface
}

func NewTransactionService(
	consumerLimitRepository repository.ConsumerLimitRepositoryInterface,
	transactionRepository repository.TransactionRepositoryInterface,
	transactionDetailRepository repository.TransactionDetailRepositoryInterface,
) *TransactionService {
	return &TransactionService{
		consumerLimitRepository:     consumerLimitRepository,
		transactionRepository:       transactionRepository,
		transactionDetailRepository: transactionDetailRepository,
	}
}

func (s *TransactionService) CreateTransaction(request *model.NewTransaction) *entity.Transaction {
	// Lock consumer limit
	key := fmt.Sprintf("%d-%d", request.ConsumerID, request.LimitID)
	m := helper.GetOrCreateMutex(key)

	m.Lock()
	defer m.Unlock()

	// Check if consumer has limit
	limit := s.consumerLimitRepository.GetByColumn(&entity.ConsumerLimit{ConsumerID: request.ConsumerID, LimitID: request.LimitID})
	if limit == nil {
		panic(utils.NotFoundError("Consumer limit not found", nil))
	}

	// Check if consumer has enough balance
	if limit.Balance < request.Amount {
		panic(utils.ClientError("Insufficient balance", nil))
	}

	// Check if consumer has reached limit
	if limit.Limit.LimitType == "percentage" {
		if limit.Balance+request.Amount > limit.Limit.LimitAmount*limit.Consumer.Salary {
			panic(utils.ClientError("Transaction amount exceeds limit", nil))
		}
	} else {
		if limit.Balance+request.Amount > limit.Limit.LimitAmount {
			panic(utils.ClientError("Transaction amount exceeds limit", nil))
		}
	}

	// Start transaction
	tx := s.transactionRepository.BeginTransaction()
	defer s.transactionRepository.CommitTransaction(tx)

	// Create transaction
	var (
		adminFee          float64 = 1000
		installmentAmount float64
		interestAmount    float64
	)

	if limit.Limit.Interest > 0 {
		interestAmount = request.Amount * limit.Limit.Interest / 100 * float64(limit.Limit.Duration)
		installmentAmount = request.Amount + adminFee + interestAmount
	} else {
		installmentAmount = request.Amount + adminFee
	}

	newTransaction := &entity.Transaction{
		ConsumerID:        request.ConsumerID,
		ContractNumber:    "123",
		OTR:               request.Amount,
		AdminFee:          adminFee,
		InstallmentAmount: installmentAmount,
		InterestAmount:    interestAmount,
		AssetName:         request.AssetName,
	}

	transaction := s.transactionRepository.WithTransaction(tx).Store(newTransaction)

	// Create transaction detail
	for i := 0; i < limit.Limit.Duration; i++ {
		newTransactionDetail := &entity.TransactionDetail{
			TransactionID:     transaction.ID,
			InstallmentAmount: installmentAmount / float64(limit.Limit.Duration),
		}

		s.transactionDetailRepository.WithTransaction(tx).Store(newTransactionDetail)
	}

	// Update consumer limit
	limit.Balance -= request.Amount
	s.consumerLimitRepository.WithTransaction(tx).UpdateByColumn(limit)

	result := s.transactionRepository.WithTransaction(tx).GetByColumn(&entity.Transaction{ID: transaction.ID})

	return result
}
