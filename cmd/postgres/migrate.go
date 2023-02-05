package postgres

import (
	"github.com/spf13/cobra"

	cmdHelper "github.com/forfam/authentication-service/cmd/helpers"
	_ "github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/src/utils/logger"
)

var log *logger.Logger

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
}
