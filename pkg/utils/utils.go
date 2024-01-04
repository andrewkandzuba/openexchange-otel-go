package utils

import "os"

const ENV_PORT = "PORT"

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
