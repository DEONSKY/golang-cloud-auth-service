package main

import (
	"github.com/spf13/cobra"

	postgresCommands "github.com/forfam/authentication-service/cmd/postgres"
	"github.com/forfam/authentication-service/src/utils/logger"
)

var log *logger.Logger

var mainCommand = &cobra.Command{
	Use:   "",
	Short: "Main command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	log = logger.New("AUTHENTICATION_SERVICE", "CommandRunner")
}

func main() {
	mainCommand.AddCommand(postgresCommands.MigrateCommand)

	if err := mainCommand.Execute(); err != nil {
		log.Fatal("Something went wrong while running command!")
	}

	log.Info("Command runned successfully!")
}
