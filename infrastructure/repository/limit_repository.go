package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/utils"
	"github.com/polluxdev/trx-core-svc/domain/entity"
)

type LimitRepositoryInterface interface {
	Store(*entity.Limit) *entity.Limit
	GetList(*string, []interface{}, map[string]interface{}) ([]*entity.Limit, *int64)
	GetByColumn(*entity.Limit) *entity.Limit
	UpdateByColumn(*entity.Limit)
	DeleteById(int)

	WithTransaction(*gorm.DB) LimitRepositoryInterface
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB)
}

type LimitRepository struct {
	dbConn *gorm.DB
}

func NewLimitRepository(dbConn *gorm.DB) LimitRepositoryInterface {
	return &LimitRepository{dbConn: dbConn}
}

func (p *LimitRepository) WithTransaction(tx *gorm.DB) LimitRepositoryInterface {
	return &LimitRepository{dbConn: tx}
}

func (p *LimitRepository) BeginTransaction() *gorm.DB {
	tx := p.dbConn.Begin()
	if tx.Error != nil {
		panic(utils.InvariantError("Failed to save data", tx.Error))
	}

	return tx
}

func (p *LimitRepository) CommitTransaction(tx *gorm.DB) {
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(utils.InvariantError("Failed to save data", err))
	}
}

func (p *LimitRepository) Store(data *entity.Limit) *entity.Limit {
	if err := p.dbConn.Create(data).Error; err != nil {
		panic(utils.InvariantError("Failed to create new data", err))
	}

	return data
}

func (p *LimitRepository) GetList(conditions *string, args []interface{}, pagination map[string]interface{}) ([]*entity.Limit, *int64) {
	var (
		result []*entity.Limit
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

func (p *LimitRepository) GetByColumn(condition *entity.Limit) *entity.Limit {
	var result entity.Limit
	if err := p.dbConn.Preload("SubConsumers").Where(condition).First(&result).Error; err != nil {
		panic(utils.InvariantError(global.DATA_FETCH_FAILED, err))
	}

	return &result
}

func (p *LimitRepository) UpdateByColumn(model *entity.Limit) {
	if err := p.dbConn.Model(model).Updates(model).Error; err != nil {
		panic(utils.InvariantError("Failed to update data", err))
	}
}

func (p *LimitRepository) DeleteById(id int) {
	if err := p.dbConn.Where("id = ?", id).Delete(&entity.Limit{}).Error; err != nil {
		panic(utils.InvariantError("Failed to delete data", err))
	}
}
