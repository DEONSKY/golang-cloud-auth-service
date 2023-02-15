package postgres

import (
	"github.com/forfam/authentication-service/config"
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
