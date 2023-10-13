package config

import (
	"os"

	_ "github.com/mileapp_screening/utils/env"
)

func init() {
	db := NewDbConfig()
	db.Load()
}

func Env() string {
	return os.Getenv("ENV")
}
