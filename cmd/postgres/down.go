package postgres

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/gorm"

	cmdHelper "github.com/forfam/authentication-service/cmd/helpers"
	"github.com/forfam/authentication-service/postgres"
)

var MigrateUndoCommand = &cobra.Command{
	Use:   "psql:migrate-undo",
	Short: "Undo migrations via gorm",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmdHelper.ParseFlag(cmd, "name", true)

		migration := findMigration(name)

		if migration == nil {
			log.Fatal(fmt.Sprintf(`Migration "%s" not found!`, name))
		}

		db := postgres.New(
			postgres.GetAuthenticationDbConfig(),
			&gorm.Config{},
		)

		isExecuted := findExecutedMigration(db, migration.Name)

		if isExecuted {

			transaction := db.Begin()

			if err := migration.Down(transaction); err != nil {
				transaction.Rollback()
				log.Fatal(fmt.Sprintf(`Something went wrong due "%s" undo migration`, migration.Name, err))
			}

			if err := unmarkMigrationMigrated(transaction, migration.Name); err != nil {
				transaction.Rollback()
				log.Fatal(fmt.Sprintf(`Something went wrong due "%s" undo migration`, migration.Name, err))
			}
			transaction.Commit()

		} else {
			log.Warning(fmt.Sprintf(`Migration: "%s" not executed before`, migration.Name))
		}
	},
}

func init() {
	MigrateUndoCommand.Flags().StringP("name", "n", "", "Name of migration to run")
}
