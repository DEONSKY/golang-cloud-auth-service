package migrations

import (
	"github.com/forfam/authentication-service/src/model"
	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230129225715",
		Up:      mig_20230129225715_init_schema_up,
		Down:    mig_20230129225715_init_schema_down,
	})
}

func mig_20230129225715_init_schema_up(db *gorm.DB) error {

	db.Migrator().CreateTable(&model.User{})
	return nil
}

func mig_20230129225715_init_schema_down(db *gorm.DB) error {
	db.Migrator().DropTable(&model.User{})
	return nil
}