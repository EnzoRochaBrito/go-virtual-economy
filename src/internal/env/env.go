package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	godotenv.Load()
}

func Get(k string, fallback string) string {
	val := os.Getenv(k)
	if val == "" {
		return fallback
	}
	return val
}
