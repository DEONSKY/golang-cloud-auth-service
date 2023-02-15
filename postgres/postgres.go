package postgres

import (
	"gorm.io/gorm"

	"github.com/forfam/authentication-service/log"
)

var AuthenticationDb *gorm.DB

var logger *log.Logger

func InitAuthenticationDb() {
	AuthenticationDb = New(
		GetAuthenticationDbConfig(),
		&gorm.Config{},
	)
}

func init() {
	logger = log.New("PostgresModule")
}
