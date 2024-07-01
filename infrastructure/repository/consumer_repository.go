package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/utils"
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type ConsumerRepositoryInterface interface {
	Store(*entity.Consumer) *entity.Consumer
	GetList(*string, []interface{}, map[string]interface{}) ([]*entity.Consumer, *int64)
	GetByColumn(*entity.Consumer) *entity.Consumer
	UpdateByColumn(*entity.Consumer)
	DeleteById(int)

	WithTransaction(*gorm.DB) ConsumerRepositoryInterface
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB)
}

type ConsumerRepository struct {
	dbConn *gorm.DB
}

func NewConsumerRepository(dbConn *gorm.DB) ConsumerRepositoryInterface {
	return &ConsumerRepository{dbConn: dbConn}
}

func (p *ConsumerRepository) WithTransaction(tx *gorm.DB) ConsumerRepositoryInterface {
	return &ConsumerRepository{dbConn: tx}
}

func (p *ConsumerRepository) BeginTransaction() *gorm.DB {
	tx := p.dbConn.Begin()
	if tx.Error != nil {
		panic(utils.InvariantError("Failed to save data", tx.Error))
	}

	return tx
}

func (p *ConsumerRepository) CommitTransaction(tx *gorm.DB) {
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(utils.InvariantError("Failed to save data", err))
	}
}

func (p *ConsumerRepository) Store(data *entity.Consumer) *entity.Consumer {
	if err := p.dbConn.Create(data).Error; err != nil {
		panic(utils.InvariantError("Failed to create new data", err))
	}

	return data
}

func (p *ConsumerRepository) GetList(conditions *string, args []interface{}, pagination map[string]interface{}) ([]*entity.Consumer, *int64) {
	var (
		result []*entity.Consumer
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

func (p *ConsumerRepository) GetByColumn(condition *entity.Consumer) *entity.Consumer {
	var result entity.Consumer
	if err := p.dbConn.Preload("SubConsumers").Where(condition).First(&result).Error; err != nil {
		panic(utils.InvariantError(global.DATA_FETCH_FAILED, err))
	}

	return &result
}

func (p *ConsumerRepository) UpdateByColumn(model *entity.Consumer) {
	if err := p.dbConn.Model(model).Updates(model).Error; err != nil {
		panic(utils.InvariantError("Failed to update data", err))
	}
}

func (p *ConsumerRepository) DeleteById(id int) {
	if err := p.dbConn.Where("id = ?", id).Delete(&entity.Consumer{}).Error; err != nil {
		panic(utils.InvariantError("Failed to delete data", err))
	}
}
