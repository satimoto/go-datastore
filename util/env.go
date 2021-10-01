package util

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func ParseBool(value string, fallback bool) bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		b = fallback
	}
	return b
}
