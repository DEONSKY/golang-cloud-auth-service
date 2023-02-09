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
		db := postgres.New(
			postgres.GetAuthenticationDbConfig(),
			&gorm.Config{},
		)

		migration := findMigration(name)

		if migration == nil {
			log.Fatal(fmt.Sprintf(`Migration "%s" not found!`, name))
		}

		transaction := db.Begin()

		if err := migration.Down(transaction); err != nil {
			transaction.Rollback()
			log.Fatal(fmt.Sprintf(`Something went wrong due "%s" removing migration`, migration.Name, err))
		}

		if err := unmarkMigrationUndone(transaction, migration.Name); err != nil {
			transaction.Rollback()
			log.Fatal(fmt.Sprintf(`Something went wrong due "%s" removing migration`, migration.Name, err))
		}
		transaction.Commit()
	},
}

func init() {
	MigrateUndoCommand.Flags().StringP("name", "n", "", "Name of migration to run")
}
