package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/forfam/authentication-service/src/utils/logger"
	"github.com/joho/godotenv"
)

type goenv string

const (
	GO_ENV_DEV     goenv = "development"
	GO_ENV_STAGING       = "staging"
	GO_ENV_PROD          = "production"
)

func (env goenv) String() string {
	return string(env)
}

func parseEnv(env string) (parsed goenv, err error) {
	envlist := map[goenv]struct{}{
		GO_ENV_DEV:     {},
		GO_ENV_STAGING: {},
		GO_ENV_PROD:    {},
	}

	parsed = goenv(env)
	_, ok := envlist[parsed]

	if !ok {
		return GO_ENV_DEV, errors.New(`Incompatible value for "GO_ENV"`)
	}

	return parsed, nil
}

var ENV goenv
var HTTP_PORT int

func init() {
	var err error

	err = godotenv.Load(".env")

	if err != nil {
		logger.GlobalLogger.Fatal("Error loading .env file")
	}

	logger.GlobalLogger.Info("Env variable load started...")

	ENV, err = parseEnv(os.Getenv("GO_ENV"))

	if err != nil {
		logger.GlobalLogger.Fatal(`Missing env variable "GO_ENV"`)
	}

	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))

	if err != nil {
		logger.GlobalLogger.Fatal(`Missin env variable "HTTP_PORT"`)
	}

	logger.GlobalLogger.Info("Env variable loaded successfuly")
}
