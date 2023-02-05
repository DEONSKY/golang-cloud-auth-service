package migrations

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"sort"
	"time"

	"github.com/forfam/authentication-service/src/utils/logger"
	"gorm.io/gorm"
)

type SchemaMigrations struct {
	Version string `gorm:"primaryKey;type:varchar(255)"`
}

type Migration struct {
	Version string
	Up      func(*gorm.DB) error
	Down    func(*gorm.DB) error

	done bool
}

type Migrator struct {
	db         *gorm.DB
	Migrations map[string]*Migration
}

var MigratorInstance = &Migrator{
	Migrations: map[string]*Migration{},
}

func (m *Migrator) AddMigration(mg *Migration) {
	// Add the migration to the hash with version as key
	m.Migrations[mg.Version] = mg
}

func Create(name string) error {
	time := time.Now().Format("20060102150405")

	version := time + "_" + name

	in := struct {
		Version string
	}{
		Version: version,
	}

	var out bytes.Buffer

	t := template.Must(template.ParseFiles("./data/migrations/postgres/template.txt"))
	err := t.Execute(&out, in)
	if err != nil {
		return errors.New("Unable to execute template:" + err.Error())
	}

	f, err := os.Create(fmt.Sprintf("./data/migrations/postgres/%s.go", version))
	if err != nil {
		return errors.New("Unable to create migration file:" + err.Error())
	}
	defer f.Close()

	if _, err := f.WriteString(out.String()); err != nil {
		return errors.New("Unable to write to migration file:" + err.Error())
	}

	logger.GlobalLogger.Info(fmt.Sprintf("Generated new migration file with name: %s", version))
	return nil
}

func Init(db *gorm.DB) *Migrator {
	MigratorInstance.db = db

	// Create `schema_migrations` table to remember which migrations were executed.
	tableExists := db.Migrator().HasTable(&SchemaMigrations{})

	if !tableExists {
		db.Migrator().CreateTable(&SchemaMigrations{})
	}

	var schemaMigrations []SchemaMigrations
	// Find out all the executed migrations
	db.Model(&SchemaMigrations{}).Find(&schemaMigrations)

	// Mark the migrations as Done if it is already executed
	for _, db_migration := range schemaMigrations {

		if MigratorInstance.Migrations[db_migration.Version] != nil {
			MigratorInstance.Migrations[db_migration.Version].done = true
		}
	}

	return MigratorInstance
}

func (m *Migrator) Up(step int) error {

	tx := m.db.Begin()

	count := 0

	keys := getSortedMigrationKeys(m.Migrations)

	for _, key := range keys {
		if step > 0 && count == step {
			break
		}

		mg := m.Migrations[key]

		if mg.done {
			continue
		}

		logger.GlobalLogger.Info("Running migration " + mg.Version)
		if err := mg.Up(tx); err != nil {
			tx.Rollback()
			return err
		}

		schemaMigration := SchemaMigrations{Version: mg.Version}

		if err := tx.Save(&schemaMigration).Error; err != nil {
			logger.GlobalLogger.Error(fmt.Sprintf("Error %s", mg.Version))
			tx.Rollback()
			return err
		}
		logger.GlobalLogger.Info("Finished running migration " + mg.Version)

		count++

	}

	tx.Commit()

	return nil
}

func (m *Migrator) Down(step int) error {
	tx := m.db.Begin()

	count := 0

	keys := getSortedMigrationKeys(m.Migrations)

	for i := len(keys) - 1; i >= 0; i-- {
		if step > 0 && count == step {
			break
		}

		mg := m.Migrations[keys[i]]

		if !mg.done {
			continue
		}

		logger.GlobalLogger.Info("Reverting Migration " + mg.Version)
		if err := mg.Down(tx); err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Where("version = ?", mg.Version).Delete(&SchemaMigrations{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		logger.GlobalLogger.Info("Finished reverting migration " + mg.Version)

		count++
	}

	tx.Commit()

	return nil
}

func (m *Migrator) MigrationStatus() error {

	keys := getSortedMigrationKeys(m.Migrations)
	for _, key := range keys {

		mg := m.Migrations[key]

		if mg.done {
			logger.GlobalLogger.Info(fmt.Sprintf("Migration %s... completed", key))
		} else {
			logger.GlobalLogger.Info(fmt.Sprintf("Migration %s... pending", key))
		}
	}

	return nil
}

func getSortedMigrationKeys(migrations map[string]*Migration) []string {
	keys := []string{}
	for key := range migrations {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}
