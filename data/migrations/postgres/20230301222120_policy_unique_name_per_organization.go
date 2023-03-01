package migrations

import (
	"gorm.io/gorm"
)

//Partial index is postgres native feature. With this migration we control uniquenes only for undeleted policies. It solves soft delete uniquenes problem

func mig_20230301222120_policy_unique_name_per_organization_up(transaction *gorm.DB) error {
	return transaction.Exec(`CREATE UNIQUE INDEX idx_unique_name_per_organization ON policies(name,organization_id) WHERE deleted_at IS NULL`).Error
}

func mig_20230301222120_policy_unique_name_per_organization_down(transaction *gorm.DB) error {
	return transaction.Exec(`DROP UNIQUE INDEX idx_unique_name_per_organization ON policies(name,organization_id) WHERE deleted_at IS NULL`).Error
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230301222120_policy_unique_name_per_organization",
		Up:   mig_20230301222120_policy_unique_name_per_organization_up,
		Down: mig_20230301222120_policy_unique_name_per_organization_down,
	})
}
