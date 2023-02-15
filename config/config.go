package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/forfam/authentication-service/log"
	"github.com/joho/godotenv"
)

var logger *log.Logger

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

var (
	ENV       goenv
	HTTP_PORT int
)

func GetConfig(key string, required bool) string {
	val := os.Getenv(key)

	if required == true && len(val) == 0 {
		logger.Fatal(fmt.Sprintf(`Missing env variable "%s"`, key))
	}

	return val
}

func GetConfigInt(key string, required bool) int {
	val := GetConfig(key, required)

	converted, err := strconv.Atoi(val)

	if err != nil {
		logger.Fatal(fmt.Sprintf(`Incompatible env value for "%s" should be int`, key))
	}

	return converted
}

func init() {
	logger = log.New("ConfigModule")
	var err error

	err = godotenv.Load(".env")

	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	logger.Info("Env variable load started...")

	ENV, err = parseEnv(os.Getenv("GO_ENV"))

	if err != nil {
		logger.Fatal(`Missing env variable "GO_ENV"`)
	}

	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))

	if err != nil {
		logger.Fatal(`Missing env variable "HTTP_PORT"`)
	}

	logger.Info("Env variable loaded successfully")
}
