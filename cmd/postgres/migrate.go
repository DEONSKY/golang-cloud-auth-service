package postgres

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/gorm"

	cmdHelper "github.com/forfam/authentication-service/cmd/helpers"
	migrations "github.com/forfam/authentication-service/data/migrations/postgres"
	"github.com/forfam/authentication-service/postgres"
)

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
		for _, migration := range migrations.Migrations {
			isExecuted := isMigrationExecuted(executeds, migration.Name)
			if isExecuted == false {
				if len(name) > 0 {
					if name == migration.Name {
						if err := migration.Up(transaction); err != nil {
							transaction.Rollback()
							log.Fatal(fmt.Sprintf(`Something went wrong due "%s" migration`, migration.Name, err))
						}

						if err := markMigrationMigrated(transaction, migration.Name); err != nil {
							transaction.Rollback()
							log.Fatal(fmt.Sprintf(`Something went wrong due "%s" migration`, migration.Name, err))
						}
					} else {
						return
					}
				} else {
					if err := migration.Up(transaction); err != nil {
						transaction.Rollback()
						log.Fatal(fmt.Sprintf(`Something went wrong due "%s" migration`, migration.Name, err))
					}
					if err := markMigrationMigrated(transaction, migration.Name); err != nil {
						transaction.Rollback()
						log.Fatal(fmt.Sprintf(`Something went wrong due "%s" migration`, migration.Name, err))
					}
				}
			}
		}

		transaction.Commit()
	},
}

func init() {
	MigrateCommand.Flags().StringP("name", "n", "", "Name of migration to run")
}
