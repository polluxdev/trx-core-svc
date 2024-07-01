package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	AppEnv   string
	AppName  string
	AppPort  string
	AppDebug bool
}

type Database struct {
	Host              string
	User              string
	Password          string
	DatabaseName      string
	Port              int
	DriverName        string
	MaxConnectionOpen int
	MaxConnectionIdle int
	Timezone          string
	MaxRetries        int
}

var (
	App = AppConfig{}
	Db  = Database{}
)

func Init() {
	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "" {
		loadEnvFile()
	}

	loadAppEnv()
	loadDbEnv()
}

func loadEnvFile() {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	if err := godotenv.Load(exPath + "/.env"); err != nil {
		logrus.WithFields(logrus.Fields{
			"executable":  ex,
			"filepath":    exPath,
			"environment": os.Getenv("APP_ENV"),
			"error":       err.Error(),
		}).Fatalln(".env is not loaded properly")
		os.Exit(1)
	}
}

func loadAppEnv() {
	App.AppEnv = os.Getenv("APP_ENV")
	App.AppName = os.Getenv("APP_NAME")
	App.AppPort = os.Getenv("APP_PORT")
	App.AppDebug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
}

func loadDbEnv() {
	Db.Host = os.Getenv("DB_HOST")
	Db.User = os.Getenv("DB_USER")
	Db.Password = os.Getenv("DB_PASSWORD")
	Db.DatabaseName = os.Getenv("DB_NAME")
	Db.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	Db.DriverName = os.Getenv("DB_DRIVER_NAME")
	Db.MaxConnectionOpen, _ = strconv.Atoi(os.Getenv("DB_MAX_CONNECTION_OPEN"))
	Db.MaxConnectionIdle, _ = strconv.Atoi(os.Getenv("DB_MAX_CONNECTION_IDLE"))
	Db.Timezone = os.Getenv("DB_TIMEZONE")
	Db.MaxRetries, _ = strconv.Atoi(os.Getenv("DB_MAX_RETRIES"))
}
