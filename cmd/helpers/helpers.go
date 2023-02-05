package cmdHelper

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/forfam/authentication-service/src/utils/logger"
)

var log *logger.Logger

func ParseFlag(cmd *cobra.Command, key string, isOptional bool) (string, error) {
	val, err := cmd.Flags().GetString(key)
	if err != nil && isOptional == false {
		log.Fatal(fmt.Sprintf(`Something went wrong while parsing "%s" flag. `, key, err))
	}
	return val, err
}

func ParseIntFlag(cmd *cobra.Command, key string, isOptional bool) (int, error) {
	val, err := ParseFlag(cmd, key, isOptional)
	if err != nil && isOptional == true {
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
