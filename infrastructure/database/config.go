package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/polluxdev/trx-core-svc/application/config"
	"github.com/sirupsen/logrus"
)

type DbConnection struct {
	Db *gorm.DB
}

func attemptDBConnection(attempt int, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(config.Db.DriverName, dsn)
	if err != nil {
		logrus.Warningf("MySQL Connect Attempt(%d) %v : %v...\n", attempt, err, dsn)
		return nil, err
	}

	logrus.Infoln("Connected to MySQL Database Successfully!")
	return db, nil
}

func setupDB(db *gorm.DB) *DbConnection {
	if config.App.AppEnv != "production" {
		db.LogMode(true)
	}

	if config.App.AppEnv == "docker" {
		db.AutoMigrate(&Consumer{}, &Limit{}, &ConsumerLimit{}, &Transaction{}, &TransactionDetail{})
	}

	db.DB().SetMaxIdleConns(config.Db.MaxConnectionIdle)
	db.DB().SetMaxOpenConns(config.Db.MaxConnectionOpen)

	if config.App.AppDebug {
		db = db.Debug()
	}

	// Seed data
	seedData(db)

	return &DbConnection{Db: db}
}

func NewConnection() *DbConnection {
	var (
		db  *gorm.DB
		err error
	)

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Db.User, config.Db.Password, config.Db.Host, config.Db.DatabaseName)
	db, err = attemptDBConnection(0, dsn)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"db_name":    config.Db.DatabaseName,
			"driver":     config.Db.DriverName,
			"connection": dsn,
			"error":      err.Error(),
		}).Warningln("Failed to connect MySQL database. Trying Reconnecting...")
	} else {
		return setupDB(db)
	}

	maxRetries := config.Db.MaxRetries
	for i := 1; i <= maxRetries; i++ {
		time.Sleep(3 * time.Second)

		db, err = attemptDBConnection(i, dsn)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"db_name":    config.Db.DatabaseName,
				"driver":     config.Db.DriverName,
				"connection": dsn,
				"error":      err.Error(),
			}).Warningln("Failed to connect MySQL database. Trying Reconnecting...")
		} else {
			return setupDB(db)
		}

		if i == maxRetries {
			logrus.WithFields(logrus.Fields{
				"db_name":    config.Db.DatabaseName,
				"driver":     config.Db.DriverName,
				"connection": dsn,
				"error":      err.Error(),
			}).Fatalln("MySQL Database Connection Failed!")
		}
	}

	return nil
}
