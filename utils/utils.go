package utils

import (
	"log"
	"os"
)

func GetEnv(path string) string {
	val := os.Getenv(path)
	if val == "" {
		log.Fatalf("Env variable for path: %s is empty.", path)
	}
	return val
}

func Contains(arr []string, seek string) bool {
	for _, val := range arr {
		if val == seek {
			return true
		}
	}
	return false
}
