package use_case

import "os"

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("No " + key)
	}
	return value
}
