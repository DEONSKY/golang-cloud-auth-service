package postgres

import "github.com/forfam/authentication-service/src/utils/logger"

var log *logger.Logger

func init() {
	log = logger.New("AUTHENTICATION_SERVICE", "PostgresMigrationCommands")
}
