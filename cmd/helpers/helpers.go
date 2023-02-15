package helpers

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/forfam/authentication-service/src/utils/logger"
)

var log *logger.Logger

func ParseFlag(cmd *cobra.Command, key string, isRequired bool) (string, error) {
	val, err := cmd.Flags().GetString(key)
	if err != nil && isRequired == true {
		log.Fatal(fmt.Sprintf(`Something went wrong while parsing "%s" flag. `, key, err))
	} else if len(val) == 0 && isRequired == true {
		log.Fatal(fmt.Sprintf(`Missing parameter "%s"!`, key))
	}

	return val, err
}

func ParseIntFlag(cmd *cobra.Command, key string, isRequired bool) (int, error) {
	val, err := ParseFlag(cmd, key, isRequired)
	if err != nil && isRequired == false {
		return 0, nil
	}

	converted, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(fmt.Sprintf(`"%s" flag is not a number`, key))
	}

	return converted, err
}

func init() {
	log = logger.New("AUTHENTICATION_SERVICE", "CommandHelper")
}
