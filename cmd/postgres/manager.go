package postgres

import (
	"sort"

	"gorm.io/gorm"

	migrations "github.com/forfam/authentication-service/data/migrations/postgres"
)

type MigrationSchema struct {
	Name string `gorm:"primaryKey;type:varchar(255)"`
}

func createTable(db *gorm.DB) {
	if exists := db.Migrator().HasTable(&MigrationSchema{}); exists == false {
		db.Migrator().CreateTable(&MigrationSchema{})
	}
}

func markMigrationMigrated(transaction *gorm.DB, name string) error {
	return transaction.Create(&MigrationSchema{name}).Error
}

func unmarkMigrationMigrated(transaction *gorm.DB, name string) error {
	return transaction.Delete(&MigrationSchema{name}).Error
}

func findMigration(name string) *migrations.PostgresMigration {
	for _, v := range migrations.Migrations {
		if v.Name == name {
			return &v
		}
	}

	return nil
}

func isMigrationExecuted(executeds []MigrationSchema, name string) bool {
	for _, v := range executeds {
		if v.Name == name {
			return true
		}
	}

	return false
}

func GetExecutedMigrations(db *gorm.DB) []MigrationSchema {
	createTable(db)
	var executedMigrations []MigrationSchema
	db.Model(&MigrationSchema{}).Find(&executedMigrations)
	SortMigrationSchemas(executedMigrations)
	return executedMigrations
}

type ByName []MigrationSchema

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func SortMigrationSchemas(migrations []MigrationSchema) {
	sort.Sort(ByName(migrations))
}
