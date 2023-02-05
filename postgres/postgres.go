package postgres

import (
	"gorm.io/gorm"
)

var AuthenticationDb *gorm.DB

func init() {
	AuthenticationDb = New(
		GetAuthenticationDbConfig(),
		&gorm.Config{},
	)
}
