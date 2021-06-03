package vault

import (
	"os"
)

func GetEnv() string {
	e := os.Getenv("ENV")

	return e
}
