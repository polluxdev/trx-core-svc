package serializer

import (
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type TransactionDTO struct {
	ID                 string                  `json:"id"`
	ConsumerID         string                  `json:"consumer_id"`
	ContractNumber     string                  `json:"contract_number"`
	OTR                float64                 `json:"otr"`
	AdminFee           float64                 `json:"admin_fee"`
	InstallmentAmount  float64                 `json:"installment_amount"`
	InterestAmount     float64                 `json:"interest_amount"`
	AssetName          string                  `json:"asset_name"`
	CreatedAt          string                  `json:"created_at"`
	TransactionDetails []*TransactionDetailDTO `json:"transaction_details"`
}

type TransactionDetailDTO struct {
	ID                string  `json:"id"`
	TransactionID     string  `json:"transaction_id"`
	InstallmentAmount float64 `json:"installment_amount"`
	CreatedAt         string  `json:"created_at"`
}

func SerializeTransaction(data *entity.Transaction) *TransactionDTO {
	transactionID := helper.Encrypt(data.ID)
	result := &TransactionDTO{
		ID:                transactionID,
		ConsumerID:        helper.Encrypt(data.ConsumerID),
		ContractNumber:    data.ContractNumber,
		OTR:               data.OTR,
		AdminFee:          data.AdminFee,
		InstallmentAmount: data.InstallmentAmount,
		InterestAmount:    data.InterestAmount,
		AssetName:         data.AssetName,
		CreatedAt:         helper.TimeToString(data.CreatedAt, global.DATE_TIME_FORMAT),
	}

	transactionDetails := make([]*TransactionDetailDTO, 0)
	for _, item := range data.TransactionDetails {
		transactionDetails = append(transactionDetails, &TransactionDetailDTO{
			ID:                helper.Encrypt(item.ID),
			TransactionID:     transactionID,
			InstallmentAmount: item.InstallmentAmount,
			CreatedAt:         helper.TimeToString(item.CreatedAt, global.DATE_TIME_FORMAT),
		})
	}

	result.TransactionDetails = transactionDetails

	return result
}

func SerializeTransactions(data []*entity.Transaction) []*TransactionDTO {
	result := make([]*TransactionDTO, 0)
	for _, item := range data {
		result = append(result, SerializeTransaction(item))
	}

	return result
}
