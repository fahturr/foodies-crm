package config

import "os"

const (
	APP_PORT  = "APP_PORT"
	TOKEN_KEY = "TOKEN_KEY"
	TOKEN_EXP = "TOKEN_EXP"
)

func GetString(key string) string {
	return os.Getenv(key)
}
