package ec

import (
	"os"
)

func DEBUG() bool {
	return os.Getenv("EC_DEBUG") != ""
}
