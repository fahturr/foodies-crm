package config

import "os"

const (
	APP_PORT = "APP_PORT"
)

func GetString(key string) string {
	return os.Getenv(key)
}
