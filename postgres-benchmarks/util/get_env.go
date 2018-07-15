package util

import "os"

// GetEnv return environment variable with default fallback if not exists
func GetEnv(env string, fallback string) string {
	env, ok := os.LookupEnv(env)
	if ok {
		return env
	}

	return fallback
}
