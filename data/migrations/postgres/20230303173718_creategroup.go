package migrations

import (
	"time"

	"gorm.io/gorm"
)

type mig_20230303173718_creategroup_struct struct {
	Id             string                                `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string                                `gorm:"type:varchar(255);not null"`
	Description    string                                `gorm:"text"`
	OrganizationId string                                `gorm:"type:uuid;not null"`
	Organization   mig_20230216213520_createorganization `gorm:"foreignkey:OrganizationId;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (entity *mig_20230303173718_creategroup_struct) TableName() string {
	return "groups"
}

func mig_20230303173718_creategroup_up(transaction *gorm.DB) error {
	return transaction.Migrator().CreateTable(&mig_20230303173718_creategroup_struct{})
}

func mig_20230303173718_creategroup_down(transaction *gorm.DB) error {
	return transaction.Migrator().DropTable(&mig_20230303173718_creategroup_struct{})
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230303173718_creategroup",
		Up:   mig_20230303173718_creategroup_up,
		Down: mig_20230303173718_creategroup_down,
	})
}
