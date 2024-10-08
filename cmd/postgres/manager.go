package postgres

import (
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

func findExecutedMigration(db *gorm.DB, name string) bool {
	row := MigrationSchema{name}
	if result := db.First(&row); result.Error != nil || result.RowsAffected == 0 {
		return false
	}

	return true
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
	db.Model(&MigrationSchema{}).Order("name").Find(&executedMigrations)
	return executedMigrations
}
