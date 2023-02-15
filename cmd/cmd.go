package main

import (
	"github.com/spf13/cobra"

	postgresCommands "github.com/forfam/authentication-service/cmd/postgres"
	"github.com/forfam/authentication-service/log"
)

var logger *log.Logger

var mainCommand = &cobra.Command{
	Use:   "",
	Short: "Main command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func main() {
	logger = log.New("CommandRunner")
	mainCommand.AddCommand(
		postgresCommands.CreateMigrationCommand,
		postgresCommands.MigrateCommand,
		postgresCommands.MigrateUndoCommand,
	)
	if err := mainCommand.Execute(); err != nil {
		logger.Fatal(`Something went wrong while running command!`)
	}

	logger.Info(`Command runned successfully!`)
}
