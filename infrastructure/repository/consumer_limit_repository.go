package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/utils"
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type ConsumerLimitRepositoryInterface interface {
	Store(*entity.ConsumerLimit) *entity.ConsumerLimit
	GetList(*string, []interface{}, map[string]interface{}) ([]*entity.ConsumerLimit, *int64)
	GetByColumn(*entity.ConsumerLimit) *entity.ConsumerLimit
	UpdateByColumn(*entity.ConsumerLimit)
	DeleteById(int)

	WithTransaction(*gorm.DB) ConsumerLimitRepositoryInterface
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB)
}

type ConsumerLimitRepository struct {
	dbConn *gorm.DB
}

func NewConsumerLimitRepository(dbConn *gorm.DB) ConsumerLimitRepositoryInterface {
	return &ConsumerLimitRepository{dbConn: dbConn}
}

func (p *ConsumerLimitRepository) WithTransaction(tx *gorm.DB) ConsumerLimitRepositoryInterface {
	return &ConsumerLimitRepository{dbConn: tx}
}

func (p *ConsumerLimitRepository) BeginTransaction() *gorm.DB {
	tx := p.dbConn.Begin()
	if tx.Error != nil {
		panic(utils.InvariantError("Failed to save data", tx.Error))
	}

	return tx
}

func (p *ConsumerLimitRepository) CommitTransaction(tx *gorm.DB) {
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(utils.InvariantError("Failed to save data", err))
	}
}

func (p *ConsumerLimitRepository) Store(data *entity.ConsumerLimit) *entity.ConsumerLimit {
	if err := p.dbConn.Create(data).Error; err != nil {
		panic(utils.InvariantError("Failed to create new data", err))
	}

	return data
}

func (p *ConsumerLimitRepository) GetList(conditions *string, args []interface{}, pagination map[string]interface{}) ([]*entity.ConsumerLimit, *int64) {
	var (
		result []*entity.ConsumerLimit
		total  int64
	)
	query := p.dbConn.Model(&result).Preload("Consumer").Preload("Limit")

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

func (p *ConsumerLimitRepository) GetByColumn(condition *entity.ConsumerLimit) *entity.ConsumerLimit {
	var result entity.ConsumerLimit
	if err := p.dbConn.Preload("Consumer").Preload("Limit").Where(condition).First(&result).Error; err != nil {
		panic(utils.InvariantError(global.DATA_FETCH_FAILED, err))
	}

	return &result
}

func (p *ConsumerLimitRepository) UpdateByColumn(model *entity.ConsumerLimit) {
	if err := p.dbConn.Model(model).Updates(model).Error; err != nil {
		panic(utils.InvariantError("Failed to update data", err))
	}
}

func (p *ConsumerLimitRepository) DeleteById(id int) {
	if err := p.dbConn.Where("id = ?", id).Delete(&entity.ConsumerLimit{}).Error; err != nil {
		panic(utils.InvariantError("Failed to delete data", err))
	}
}
