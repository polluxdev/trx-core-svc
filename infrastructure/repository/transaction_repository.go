package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/utils"
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type TransactionRepositoryInterface interface {
	Store(*entity.Transaction) *entity.Transaction
	GetList(*string, []interface{}, map[string]interface{}) ([]*entity.Transaction, *int64)
	GetByColumn(*entity.Transaction) *entity.Transaction
	UpdateByColumn(*entity.Transaction)
	DeleteById(int)

	WithTransaction(*gorm.DB) TransactionRepositoryInterface
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB)
}

type TransactionRepository struct {
	dbConn *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) TransactionRepositoryInterface {
	return &TransactionRepository{dbConn: dbConn}
}

func (p *TransactionRepository) WithTransaction(tx *gorm.DB) TransactionRepositoryInterface {
	return &TransactionRepository{dbConn: tx}
}

func (p *TransactionRepository) BeginTransaction() *gorm.DB {
	tx := p.dbConn.Begin()
	if tx.Error != nil {
		panic(utils.InvariantError("Failed to save data", tx.Error))
	}

	return tx
}

func (p *TransactionRepository) CommitTransaction(tx *gorm.DB) {
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(utils.InvariantError("Failed to save data", err))
	}
}

func (p *TransactionRepository) Store(data *entity.Transaction) *entity.Transaction {
	if err := p.dbConn.Create(data).Error; err != nil {
		panic(utils.InvariantError("Failed to create new data", err))
	}

	return data
}

func (p *TransactionRepository) GetList(conditions *string, args []interface{}, pagination map[string]interface{}) ([]*entity.Transaction, *int64) {
	var (
		result []*entity.Transaction
		total  int64
	)
	query := p.dbConn.Model(&result)

	if conditions != nil {
		query = query.Where(*conditions, args...)
	}

	query.Count(&total)

	if pagination["page"] != 0 || pagination["limit"] != 0 {
		query = query.Offset((pagination["page"].(int) - 1) * pagination["limit"].(int)).Limit(pagination["limit"].(int))
	}

	if err := query.Find(&result).Error; err != nil {
		panic(utils.InvariantError(global.DATA_FETCH_FAILED, err))
	}

	return result, &total
}

func (p *TransactionRepository) GetByColumn(condition *entity.Transaction) *entity.Transaction {
	var result entity.Transaction
	if err := p.dbConn.Preload("TransactionDetails").Where(condition).First(&result).Error; err != nil {
		panic(utils.InvariantError(global.DATA_FETCH_FAILED, err))
	}

	return &result
}

func (p *TransactionRepository) UpdateByColumn(model *entity.Transaction) {
	if err := p.dbConn.Model(model).Updates(model).Error; err != nil {
		panic(utils.InvariantError("Failed to update data", err))
	}
}

func (p *TransactionRepository) DeleteById(id int) {
	if err := p.dbConn.Where("id = ?", id).Delete(&entity.Transaction{}).Error; err != nil {
		panic(utils.InvariantError("Failed to delete data", err))
	}
}
