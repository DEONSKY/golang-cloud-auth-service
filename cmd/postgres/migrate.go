package postgres

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/gorm"

	cmdHelper "github.com/forfam/authentication-service/cmd/helpers"
	migrations "github.com/forfam/authentication-service/data/migrations/postgres"
	"github.com/forfam/authentication-service/postgres"
)

func migrate(migration migrations.PostgresMigration, transaction *gorm.DB) {
	if err := migration.Up(transaction); err != nil {
		transaction.Rollback()
		log.Fatal(fmt.Sprintf(`Something went wrong due "%s" migration`, migration.Name, err))
	}

	if err := markMigrationMigrated(transaction, migration.Name); err != nil {
		transaction.Rollback()
		log.Fatal(fmt.Sprintf(`Something went wrong due "%s" migration`, migration.Name, err))
	}
}

var MigrateCommand = &cobra.Command{
	Use:   "psql:migrate",
	Short: "Run migrations via gorm",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmdHelper.ParseFlag(cmd, "name", false)
		db := postgres.New(
			postgres.GetAuthenticationDbConfig(),
			&gorm.Config{},
		)

		executeds := GetExecutedMigrations(db)

		migrations.Sort()

		transaction := db.Begin()

		if len(name) > 0 {

			migration := findMigration(name)
			if migration == nil {
				log.Warning(fmt.Sprintf(`Migration not found "%s". Please check migration name exists in declared migrations folder`, name))
				return
			}

			isExecuted := isMigrationExecuted(executeds, migration.Name)

			if !isExecuted {
				migrate(*migration, transaction)

			} else {
				log.Warning(fmt.Sprintf(`Migration: "%s" already executed`, migration.Name))
			}
		} else {
			for _, migration := range migrations.Migrations {
				isExecuted := isMigrationExecuted(executeds, migration.Name)
				if !isExecuted {
					migrate(migration, transaction)
				}
			}
		}
		transaction.Commit()

	},
}

func init() {
	MigrateCommand.Flags().StringP("name", "n", "", "Name of migration to run")
}
