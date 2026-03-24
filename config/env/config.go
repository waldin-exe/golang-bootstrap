package config

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		if v, err := strconv.Atoi(val); err == nil {
			return v
		}
	}
	return fallback
}
