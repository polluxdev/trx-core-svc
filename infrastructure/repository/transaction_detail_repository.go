package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/utils"
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type TransactionDetailRepositoryInterface interface {
	Store(*entity.TransactionDetail) *entity.TransactionDetail
	GetList(*string, []interface{}, map[string]interface{}) ([]*entity.TransactionDetail, *int64)
	GetByColumn(*entity.TransactionDetail) *entity.TransactionDetail
	UpdateByColumn(*entity.TransactionDetail)
	DeleteById(int)

	WithTransaction(*gorm.DB) TransactionDetailRepositoryInterface
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB)
}

type TransactionDetailRepository struct {
	dbConn *gorm.DB
}

func NewTransactionDetailRepository(dbConn *gorm.DB) TransactionDetailRepositoryInterface {
	return &TransactionDetailRepository{dbConn: dbConn}
}

func (p *TransactionDetailRepository) WithTransaction(tx *gorm.DB) TransactionDetailRepositoryInterface {
	return &TransactionDetailRepository{dbConn: tx}
}

func (p *TransactionDetailRepository) BeginTransaction() *gorm.DB {
	tx := p.dbConn.Begin()
	if tx.Error != nil {
		panic(utils.InvariantError("Failed to save data", tx.Error))
	}

	return tx
}

func (p *TransactionDetailRepository) CommitTransaction(tx *gorm.DB) {
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(utils.InvariantError("Failed to save data", err))
	}
}

func (p *TransactionDetailRepository) Store(data *entity.TransactionDetail) *entity.TransactionDetail {
	if err := p.dbConn.Create(data).Error; err != nil {
		panic(utils.InvariantError("Failed to create new data", err))
	}

	return data
}

func (p *TransactionDetailRepository) GetList(conditions *string, args []interface{}, pagination map[string]interface{}) ([]*entity.TransactionDetail, *int64) {
	var (
		result []*entity.TransactionDetail
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

func (p *TransactionDetailRepository) GetByColumn(condition *entity.TransactionDetail) *entity.TransactionDetail {
	var result entity.TransactionDetail
	if err := p.dbConn.Where(condition).First(&result).Error; err != nil {
		panic(utils.InvariantError(global.DATA_FETCH_FAILED, err))
	}

	return &result
}

func (p *TransactionDetailRepository) UpdateByColumn(model *entity.TransactionDetail) {
	if err := p.dbConn.Model(model).Updates(model).Error; err != nil {
		panic(utils.InvariantError("Failed to update data", err))
	}
}

func (p *TransactionDetailRepository) DeleteById(id int) {
	if err := p.dbConn.Where("id = ?", id).Delete(&entity.TransactionDetail{}).Error; err != nil {
		panic(utils.InvariantError("Failed to delete data", err))
	}
}
