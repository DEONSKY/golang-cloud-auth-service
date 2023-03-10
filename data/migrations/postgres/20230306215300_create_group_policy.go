package migrations

import (
	"time"

	"gorm.io/gorm"
)

type mig_20230306215300_create_group_policy_struct struct {
	Id        string                                 `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PolicyId  string                                 `gorm:"type:uuid;not null"`
	Policy    mig_20230227213100_createpolicy_struct `gorm:"foreignkey:PolicyId;"`
	GroupId   string                                 `gorm:"type:uuid;not null"`
	Group     mig_20230303173718_creategroup_struct  `gorm:"foreignkey:GroupId;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (entity *mig_20230306215300_create_group_policy_struct) TableName() string {
	return "group_policies"
}

func mig_20230306215300_create_group_policy_up(transaction *gorm.DB) error {
	return transaction.Migrator().CreateTable(&mig_20230306215300_create_group_policy_struct{})
}

func mig_20230306215300_create_group_policy_down(transaction *gorm.DB) error {
	return transaction.Migrator().DropTable(&mig_20230306215300_create_group_policy_struct{})
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230306215300_create_group_policy",
		Up:   mig_20230306215300_create_group_policy_up,
		Down: mig_20230306215300_create_group_policy_down,
	})
}
