package psqlmigcmd

import (
	"github.com/spf13/cobra"

	migrations "github.com/forfam/authentication-service/cmd/postgres/service"
	_ "github.com/forfam/authentication-service/postgres"
)

var migrateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {

		migrations.Init(authenticationDb)

	},
}
