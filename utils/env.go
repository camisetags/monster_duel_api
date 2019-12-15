package utils

import (
	"os"
)

// GetEnvOrDefault will get env var or default value
func GetEnvOrDefault(envVar, defaultValue string) string {
	value := os.Getenv(envVar)

	if len(value) == 0 {
		return defaultValue
	}

	return value
}
