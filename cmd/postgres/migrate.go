package psqlmigcmd

import (
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	cmdHelper "github.com/forfam/authentication-service/cmd/helpers"
	"github.com/forfam/authentication-service/postgres"

	"github.com/forfam/authentication-service/src/utils/logger"
)

var log *logger.Logger
var authenticationDb *gorm.DB

var MigrateCommand = &cobra.Command{
	Use:   "psql:migrate",
	Short: "Run migrations via gorm",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmdHelper.ParseFlag(cmd, "name", true)
		log.Info("migrate command runned!" + name)
	},
}

func init() {
	log = logger.New("AUTHENTICATION_SERVICE", "PostgresMigrationCMD")
	authenticationDb = postgres.New(
		postgres.GetAuthenticationDbConfig(),
		&gorm.Config{},
	)
	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")

	// Add "--step" flag to both "up" and "down" command
	migrateUpCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	migrateDownCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")

	// Add "create", "up" and "down" commands to the "migrate" command
	MigrateCommand.AddCommand(migrateUpCmd, migrateDownCmd, migrateCreateCmd, migrateStatusCmd)
}
