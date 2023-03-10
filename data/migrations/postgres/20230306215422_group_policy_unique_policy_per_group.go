package migrations

import (
	"gorm.io/gorm"
)

func mig_20230306215422_group_policy_unique_policy_per_group_up(transaction *gorm.DB) error {
	return transaction.Exec(`CREATE UNIQUE INDEX idx_group_policy_unique_policy_per_group ON group_policies(policy_id,group_id) WHERE deleted_at IS NULL`).Error
}

func mig_20230306215422_group_policy_unique_policy_per_group_down(transaction *gorm.DB) error {
	return transaction.Exec(`DROP INDEX idx_group_policy_unique_policy_per_group`).Error
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230306215422_group_policy_unique_policy_per_group",
		Up:   mig_20230306215422_group_policy_unique_policy_per_group_up,
		Down: mig_20230306215422_group_policy_unique_policy_per_group_down,
	})
}
