package postgres

import (
	"gorm.io/gorm"

	"github.com/forfam/authentication-service/src/utils/logger"
)

var AuthenticationDb *gorm.DB

var log *logger.Logger

func InitAuthenticationDb() {
	AuthenticationDb = New(
		GetAuthenticationDbConfig(),
		&gorm.Config{},
	)
}

func init() {
	log = logger.New("AUTHENTICATION_SERVICE", "PostgresModule")
}
