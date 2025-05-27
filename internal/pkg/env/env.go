package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvString(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Fatal(key, " was not found in .env file.")
	return ""
}

func GetEnvInt(key string) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	log.Fatal(key, " was not found in .env file.")
	return -1
}
