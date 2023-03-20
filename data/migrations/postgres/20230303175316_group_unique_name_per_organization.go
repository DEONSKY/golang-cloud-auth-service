package migrations

import (
	"gorm.io/gorm"
)

func mig_20230303175316_group_unique_name_per_organization_up(transaction *gorm.DB) error {
	return transaction.Exec(`CREATE UNIQUE INDEX idx_group_unique_name_per_organization ON groups(name,organization_id) WHERE deleted_at IS NULL`).Error
}

func mig_20230303175316_group_unique_name_per_organization_down(transaction *gorm.DB) error {
	return transaction.Exec(`DROP INDEX idx_group_unique_name_per_organization`).Error
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230303175316_group_unique_name_per_organization",
		Up:   mig_20230303175316_group_unique_name_per_organization_up,
		Down: mig_20230303175316_group_unique_name_per_organization_down,
	})
}
