package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/forfam/authentication-service/src/utils/logger"
	"github.com/joho/godotenv"
)

var log *logger.Logger

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
		log.Fatal(fmt.Sprintf(`Missing env variable "%s"`, key))
	}

	return val
}

func GetConfigInt(key string, required bool) int {
	val := GetConfig(key, required)

	converted, err := strconv.Atoi(val)

	if err != nil {
		log.Fatal(fmt.Sprintf(`Incompatible env value for "%s" should be int`, key))
	}

	return converted
}

func init() {
	log = logger.New("AUTHENTICATION_SERVICE", "ConfigModule")
	var err error

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Info("Env variable load started...")

	ENV, err = parseEnv(os.Getenv("GO_ENV"))

	if err != nil {
		log.Fatal(`Missing env variable "GO_ENV"`)
	}

	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))

	if err != nil {
		log.Fatal(`Missing env variable "HTTP_PORT"`)
	}

	log.Info("Env variable loaded successfully")
}
