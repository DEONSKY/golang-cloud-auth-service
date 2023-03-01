package postgres

import (
	"gorm.io/gorm"

	"github.com/forfam/authentication-service/log"

	gormLogger "gorm.io/gorm/logger"
)

var AuthenticationDb *gorm.DB

var logger *log.Logger

func InitAuthenticationDb() {
	AuthenticationDb = New(
		GetAuthenticationDbConfig(),
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(DeclareLogLevel()),
		},
	)
}

func init() {
	logger = log.New("PostgresModule")
}
