package group

import (
	"github.com/forfam/authentication-service/log"
)

var logger *log.Logger

func init() {
	logger = log.New("GroupModule")
}
