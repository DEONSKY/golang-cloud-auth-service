package migration

import (
	"github.com/forfam/authentication-service/src/model"
)

type Migration struct {
	Up func();
	Down func();
}

func up(db){
	db.Migrator().CreateTable(&model.User{})
}

func down(db){
	db.Migrator().DropTable(&model.User{})
}

func NewInitialMigration() *Migration {
	migration := Migration{
		Up:    up,
		Down: down,
	}

	return &migration
}
