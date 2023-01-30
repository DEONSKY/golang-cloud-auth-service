package migrations

import (
	"github.com/forfam/authentication-service/src/model"
	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230130204748",
		Up:      mig_20230130204748_user_subject_up,
		Down:    mig_20230130204748_user_subject_down,
	})
}

func mig_20230130204748_user_subject_up(db *gorm.DB) error {
	db.Migrator().CreateTable(&model.SubjectUser{})
	db.Migrator().CreateConstraint(&model.User{}, "fk_subject_users_user")
	db.Migrator().CreateConstraint(&model.Subject{}, "fk_subject_users_subject")
	return nil
}

func mig_20230130204748_user_subject_down(db *gorm.DB) error {
	db.Migrator().DropConstraint(&model.SubjectUser{}, "fk_subject_users_user")
	db.Migrator().DropConstraint(&model.SubjectUser{}, "fk_subject_users_subject")
	db.Migrator().DropTable(&model.SubjectUser{})
	return nil
}
