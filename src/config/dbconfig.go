package config

import (
	"fmt"
	"os"
	"time"

	"github.com/forfam/authentication-service/src/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB = nil
var IsDbConnected bool = false

const dbReconnectionTimeout = 30 * time.Second
const dbReconnectionRetryCount = 10

func SetupDatabaseConnection() (*gorm.DB, error) {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", dbHost, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	GormDB = db

	if err != nil {
		IsDbConnected = false
		logger.GlobalLogger.Error("DB connection attempt failed")
		return db, err
	}

	IsDbConnected = true
	logger.GlobalLogger.Info("DB connection attempt successfull")
	return db, err

}

func AutoConnectDbLoop() {
	for counter := 0; counter < dbReconnectionRetryCount; counter++ {
		autoConnectionControl()
	}
}

func autoConnectionControl() {
	defer time.Sleep(dbReconnectionTimeout)
	if GormDB == nil {
		SetupDatabaseConnection()
		return
	}
	sqlDB, err := GormDB.DB()
	if sqlDB.Ping() != nil {
		IsDbConnected = false
	}
	if err == nil && !IsDbConnected {
		SetupDatabaseConnection()
	}
}

// CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection() {
	dbSQL, err := GormDB.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
