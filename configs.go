package main

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	godotenv.Load()
	return os.Getenv(key)
}

func EnvIsTrue(key string) bool {
	return GetEnv(key) == "true"
}
