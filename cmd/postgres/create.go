package postgres

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/spf13/cobra"

	"github.com/forfam/authentication-service/cmd/helpers"
)

const templatePath = "./cmd/postgres/migration_template.go.txt"

func createMigration(name string) string {
	time := time.Now().Format("20060102150405")
	migrationName := fmt.Sprintf("%s_%s", time, name)
	cleanedMigrationName := slug.MakeLang(migrationName, "en")
	cleanedMigrationName = strings.ReplaceAll(cleanedMigrationName, "-", "_")

	var fileBuffer bytes.Buffer
	templateData := struct{ Name string }{cleanedMigrationName}
	templateRunner := template.Must(template.ParseFiles(templatePath))
	err := templateRunner.Execute(&fileBuffer, templateData)
	if err != nil {
		log.Fatal(fmt.Sprintf(`Something went wrong while try to parse template file!`, err))
	}

	file, err := os.Create(fmt.Sprintf(`./data/migrations/postgres/%s.go`, cleanedMigrationName))
	if err != nil {
		log.Fatal(fmt.Sprintf(`Something went wrong while try to create migration file!`, err))
	}

	defer file.Close()

	if _, err := file.WriteString(fileBuffer.String()); err != nil {
		log.Fatal(fmt.Sprintf(`Something went wrong while try to write migration file!`, err))
	}

	return cleanedMigrationName
}

var CreateMigrationCommand = &cobra.Command{
	Use:   "psql:migration-create",
	Short: "Create a new empty migrations file",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := helpers.ParseFlag(cmd, "name", false)

		if err != nil {
			log.Fatal(`Something went wrong while claiming requeired "name" argument`)
			return
		}

		migrationName := createMigration(name)

		log.Info(fmt.Sprintf(`Migration %s created successfully`, migrationName))
	},
}

func init() {
	CreateMigrationCommand.Flags().StringP("name", "n", "", "Name of the migration")
}
