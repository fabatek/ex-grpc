package config

import "os"

var (
	SERVER string
)

func init() {
	SERVER = GetEnv("SERVER", "localhost:8080")
}

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}
