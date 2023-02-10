```go
package migrations

import (
	"gorm.io/gorm"
)

type subject_user_20230210160630 struct {
	Id        string                 `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string                 `gorm:"type:uuid"`
	User      users_20230210153114   `gorm:"foreignkey:UserID;"`
	SubjectID string                 `gorm:"type:uuid"`
	Subject   subject_20230210155706 `gorm:"foreignkey:SubjectID;"`
}

func (table *subject_user_20230210160630) TableName() string {
	return "subject_users"
}

type user_20230210160630 struct{}

type subject_20230210160630 struct{}

func mig_20230210160630_subject_user_create_up(transaction *gorm.DB) error {
	transaction.Migrator().CreateTable(&subject_user_20230210160630{})
	transaction.Migrator().CreateConstraint(&user_20230210160630{}, "fk_subject_users_user")
	transaction.Migrator().CreateConstraint(&subject_20230210160630{}, "fk_subject_users_subject")
	return nil
}

func mig_20230210160630_subject_user_create_down(transaction *gorm.DB) error {
	transaction.Migrator().DropConstraint(&subject_user_20230210160630{}, "fk_subject_users_user")
	transaction.Migrator().DropConstraint(&subject_user_20230210160630{}, "fk_subject_users_subject")
	transaction.Migrator().DropTable(&subject_user_20230210160630{})
	return nil
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230210160630_subject_user_create",
		Up:   mig_20230210160630_subject_user_create_up,
		Down: mig_20230210160630_subject_user_create_down,
	})
}
```
