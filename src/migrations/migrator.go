package migrations

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"time"

	"gorm.io/gorm"
)

type Migration struct {
	Version string
	Up      func(*gorm.DB) error
	Down    func(*gorm.DB) error

	done bool
}

type Migrator struct {
	db         *gorm.DB
	Versions   []string
	Migrations map[string]*Migration
}

var migrator = &Migrator{
	Versions:   []string{},
	Migrations: map[string]*Migration{},
}

func (m *Migrator) AddMigration(mg *Migration) {
	// Add the migration to the hash with version as key
	m.Migrations[mg.Version] = mg

	// Insert version into versions array using insertion sort
	index := 0
	for index < len(m.Versions) {
		if m.Versions[index] > mg.Version {
			break
		}
		index++
	}

	m.Versions = append(m.Versions, mg.Version)
	copy(m.Versions[index+1:], m.Versions[index:])
	m.Versions[index] = mg.Version
}

func Create(name string) error {
	version := time.Now().Format("20060102150405")

	in := struct {
		Version string
		Name    string
	}{
		Version: version,
		Name:    name,
	}

	var out bytes.Buffer

	t := template.Must(template.ParseFiles("./src/migrations/template.txt"))
	err := t.Execute(&out, in)
	if err != nil {
		return errors.New("Unable to execute template:" + err.Error())
	}

	f, err := os.Create(fmt.Sprintf("./src/migrations/%s_%s.go", version, name))
	if err != nil {
		return errors.New("Unable to create migration file:" + err.Error())
	}
	defer f.Close()

	if _, err := f.WriteString(out.String()); err != nil {
		return errors.New("Unable to write to migration file:" + err.Error())
	}

	fmt.Println("Generated new migration files...", f.Name())
	return nil
}

type SchemaMigrations struct {
	Version string `gorm:"primaryKey;type:varchar(255)"`
}

func Init(db *gorm.DB) *Migrator {
	migrator.db = db

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

		if migrator.Migrations[db_migration.Version] != nil {
			migrator.Migrations[db_migration.Version].done = true
		}
	}

	return migrator
}

func (m *Migrator) Up(step int) error {

	tx := m.db.Begin()

	count := 0

	for _, v := range m.Versions {
		if step > 0 && count == step {
			break
		}

		mg := m.Migrations[v]

		if mg.done {
			continue
		}

		fmt.Println("Running migration", mg.Version)
		if err := mg.Up(tx); err != nil {
			tx.Rollback()
			return err
		}

		schemaMigration := SchemaMigrations{Version: mg.Version}

		if err := tx.Save(&schemaMigration).Error; err != nil {
			fmt.Println("Error", mg.Version)
			tx.Rollback()
			return err
		}
		fmt.Println("Finished running migration", mg.Version)

		count++

	}

	tx.Commit()

	return nil
}

// Code removed for brevity

func (m *Migrator) Down(step int) error {
	tx := m.db.Begin()

	count := 0
	for _, v := range reverse(m.Versions) {
		if step > 0 && count == step {
			break
		}

		mg := m.Migrations[v]

		if !mg.done {
			continue
		}

		fmt.Println("Reverting Migration", mg.Version)
		if err := mg.Down(tx); err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Where("version = ?", mg.Version).Delete(&SchemaMigrations{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		fmt.Println("Finished reverting migration", mg.Version)

		count++
	}

	tx.Commit()

	return nil
}

func reverse(arr []string) []string {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func (m *Migrator) MigrationStatus() error {
	for _, v := range m.Versions {
		mg := m.Migrations[v]

		if mg.done {
			fmt.Println(fmt.Sprintf("Migration %s... completed", v))
		} else {
			fmt.Println(fmt.Sprintf("Migration %s... pending", v))
		}
	}

	return nil
}
