package psqlmigcmd

import (
	"github.com/spf13/cobra"

	cmdHelper "github.com/forfam/authentication-service/cmd/helpers"
	migrations "github.com/forfam/authentication-service/cmd/postgres/service"
	_ "github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/src/utils/logger"
)

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step, err := cmdHelper.ParseIntFlag(cmd, "step", true)

		if err != nil {
			logger.GlobalLogger.Fatal("Something went wrong while claiming step argument")
			return
		}

		migrator := migrations.Init(authenticationDb)

		migrator.Down(step)
	},
}
