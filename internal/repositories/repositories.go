package repositories

import (
	"log"
	"prea/internal/common"
)

func GetLogger() *log.Logger {
	return common.MakeLogger("repositories")
}
