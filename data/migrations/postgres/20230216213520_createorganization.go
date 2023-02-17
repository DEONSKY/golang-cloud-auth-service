package migrations

import (
	"time"

	"gorm.io/gorm"
)

type mig_20230216213520_createorganization struct {
	Id          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (t *mig_20230216213520_createorganization) TableName() string {
	return "organizations"
}

func mig_20230216213520_createorganization_up(transaction *gorm.DB) error {
	return transaction.Migrator().CreateTable(&mig_20230216213520_createorganization{})
}

func mig_20230216213520_createorganization_down(transaction *gorm.DB) error {
	return transaction.Migrator().DropTable(&mig_20230216213520_createorganization{})
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230216213520_createorganization",
		Up:   mig_20230216213520_createorganization_up,
		Down: mig_20230216213520_createorganization_down,
	})
}
