package cmd

import (
	"fmt"

	"log"

	migrations "github.com/forfam/authentication-service/data/migrations/postgres"
	"github.com/forfam/authentication-service/src/config"
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
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Unable to read flag `name`", err.Error())
			return
		}

		if err := migrations.Create(name); err != nil {
			fmt.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db, err := config.SetupDatabaseConnection()

		migrator := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Up(step)
		if err != nil {
			fmt.Println("Unable to run `up` migrations")
			return
		}

	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db, err := config.SetupDatabaseConnection()

		migrator := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Down(step)
		if err != nil {
			fmt.Println("Unable to run `down` migrations")
			return
		}
	},
}

var migrateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := config.SetupDatabaseConnection()

		if err != nil {
			fmt.Println("Unable to fetch migration status")
		}

		migrator := migrations.Init(db)

		if err := migrator.MigrationStatus(); err != nil {
			fmt.Println("Unable to fetch migration status")
			return
		}

		return
	},
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
		log.Fatalln(err.Error())
	}
}
