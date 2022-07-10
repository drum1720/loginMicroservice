package configs

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
)

func LoadEnvConfigs() {
	_ = godotenv.Load()
}

func GetEnvCfg(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.Errorf("%s not config in .env", key)
	}
	return value, nil
}
