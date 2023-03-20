package postgres

import (
	"github.com/forfam/authentication-service/config"

	gormLogger "gorm.io/gorm/logger"
)

func GetAuthenticationDbConfig() *DbConnectionOptions {
	return &DbConnectionOptions{
		Host:                   config.GetConfig("POSTGRES_DB_HOST", true),
		Port:                   config.GetConfigInt("POSTGRES_DB_PORT", true),
		User:                   config.GetConfig("POSTGRES_DB_USER", true),
		Pwd:                    config.GetConfig("POSTGRES_DB_PWD", true),
		Database:               config.GetConfig("POSTGRES_DB_NAME", true),
		MaxRetryCount:          5,
		ConnectionRetryTimeout: 10,
	}
}

func DeclareLogLevel() gormLogger.LogLevel {
	logLevel := gormLogger.Silent
	goenv := config.GetConfig("GO_ENV", false)
	if goenv == "development" {
		logLevel = gormLogger.Info
	}
	return logLevel
}
