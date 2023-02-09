package postgres

import (
	"gorm.io/gorm"
)

var AuthenticationDb *gorm.DB

func InitAuthenticationDb() {
	AuthenticationDb = New(
		GetAuthenticationDbConfig(),
		&gorm.Config{},
	)
}
