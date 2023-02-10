package postgres

import (
	"fmt"
	"time"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnectionOptions struct {
	Host                   string
	Port                   int
	User                   string
	Pwd                    string
	Database               string
	MaxRetryCount          int
	ConnectionRetryTimeout int
}

func connectWithRetry(
	connectionUri string,
	config *gorm.Config,
	retryCount int,
	retryTimeout int,
) *gorm.DB {
	var connection *gorm.DB
	var err error

	for i := 0; i < retryCount; i++ {
		connection, err = gorm.Open(psql.Open(connectionUri), config)
		if err == nil {
			break
		}

		log.Warning(fmt.Sprintf(`Postgres DB connection retry %d failed with:`, i, err))
		time.Sleep(time.Duration(retryTimeout) * time.Second)
	}

	if err != nil {
		log.Fatal(fmt.Sprintf(`Postgres DB connection failed with %d retry...`, retryCount, err))
	}

	log.Info("Postgres DB successfully connected!")

	return connection
}

func New(options *DbConnectionOptions, config *gorm.Config) *gorm.DB {
	if options.MaxRetryCount < 0 {
		log.Fatal(`"DbConnectionOptions.MaxRetryCount" can not be negative.`)
	}

	connectionUri := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		options.User,
		options.Pwd,
		options.Host,
		options.Port,
		options.Database,
	)

	return connectWithRetry(connectionUri, config, options.MaxRetryCount, options.ConnectionRetryTimeout)
}
