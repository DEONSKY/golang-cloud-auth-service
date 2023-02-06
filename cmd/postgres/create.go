package psqlmigcmd

import (
	"github.com/spf13/cobra"

	cmdHelper "github.com/forfam/authentication-service/cmd/helpers"
	migrations "github.com/forfam/authentication-service/cmd/postgres/service"
	_ "github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/src/utils/logger"
)

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migrations file",
	Run: func(cmd *cobra.Command, args []string) {

		name, err := cmdHelper.ParseFlag(cmd, "name", false)

		if err != nil {
			logger.GlobalLogger.Fatal("Something went wrong while claiming requeired `name` argument")
			return
		}

		if err := migrations.Create(name); err != nil {
			logger.GlobalLogger.Fatal("Unable to create migration" + err.Error())
			return
		}
	},
}
