package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//EnvVar function is for read .env file
func EnvVar(key string, defaultVal string) string {
	godotenv.Load(".env")
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	fmt.Println(value)
	return value
}
