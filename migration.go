package main

import (
	"fmt"

	migrations "github.com/forfam/authentication-service/data/migrations/postgres"
	"github.com/forfam/authentication-service/src/config"
	"github.com/forfam/authentication-service/src/utils/logger"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migration",
	Short: "Migration tool for versioned gorm migrations by",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migrations file",
	Run: func(cmd *cobra.Command, args []string) {

		name := readString(cmd, "name")

		if err := migrations.Create(name); err != nil {
			logger.GlobalLogger.Fatal("Unable to create migration" + err.Error())
			return
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step := readInt(cmd, "step")

		db, err := config.SetupDatabaseConnection()
		if err != nil {
			logger.GlobalLogger.Fatal("Something went wrong while database connection")
			return
		}

		migrator := migrations.Init(db)
		if err != nil {
			logger.GlobalLogger.Fatal("Unable to fetch migrator")
			return
		}

		migrator.Up(step)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step := readInt(cmd, "step")

		db, err := config.SetupDatabaseConnection()

		if err != nil {
			logger.GlobalLogger.Fatal("Something went wrong while database connection")
			return
		}

		migrator := migrations.Init(db)
		if err != nil {
			logger.GlobalLogger.Fatal("Unable to fetch migrator")
			return
		}

		migrator.Down(step)
	},
}

var migrateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := config.SetupDatabaseConnection()

		if err != nil {
			logger.GlobalLogger.Fatal("Unable to fetch migration status")
		}

		migrator := migrations.Init(db)

		if err := migrator.MigrationStatus(); err != nil {
			logger.GlobalLogger.Fatal("Unable to fetch migration status")
			return
		}

		return
	},
}

func readInt(cmd *cobra.Command, key string) int {
	val, err := cmd.Flags().GetInt(key)
	if err != nil {
		logger.GlobalLogger.Fatal(fmt.Sprintf("Missing flag `%s`", key))
	}
	return val
}

func readString(cmd *cobra.Command, key string) string {
	val, err := cmd.Flags().GetString(key)
	if err != nil {
		logger.GlobalLogger.Fatal(fmt.Sprintf("Missing flag `%s`", key))
	}
	return val
}

func init() {
	// Add "--name" flag to "create" command
	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")

	// Add "--step" flag to both "up" and "down" command
	migrateUpCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	migrateDownCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")

	// Add "create", "up" and "down" commands to the "migrate" command
	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateCreateCmd, migrateStatusCmd)

}

func ExecuteMigrationTool() {
	if err := migrateCmd.Execute(); err != nil {
		logger.GlobalLogger.Fatal(err.Error())
	}
}

func main() {
	ExecuteMigrationTool()
}
