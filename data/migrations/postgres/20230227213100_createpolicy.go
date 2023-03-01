package migrations

import (
	"time"

	"gorm.io/gorm"
)

type mig_20230227213100_createpolicy_struct struct {
	Id             string                                `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string                                `gorm:"type:varchar(255);index:idx_unique_name_per_organization,unique;not null"`
	OrganizationId string                                `gorm:"type:uuid;index:idx_unique_name_per_organization,unique;not null"`
	Organization   mig_20230216213520_createorganization `gorm:"foreignkey:OrganizationId;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (entity *mig_20230227213100_createpolicy_struct) TableName() string {
	return "policies"
}

func mig_20230227213100_createpolicy_up(transaction *gorm.DB) error {
	return transaction.Migrator().CreateTable(&mig_20230227213100_createpolicy_struct{})
}

func mig_20230227213100_createpolicy_down(transaction *gorm.DB) error {
	return transaction.Migrator().DropTable(&mig_20230227213100_createpolicy_struct{})
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230227213100_createpolicy",
		Up:   mig_20230227213100_createpolicy_up,
		Down: mig_20230227213100_createpolicy_down,
	})
}
