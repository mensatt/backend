package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	return v
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panic("ENV err: [" + k + "]" + err.Error())
	}
	return b
}

// MustGetInt32 will return the env as int32 or panic if it is not present
func MustGetInt32(k string) int32 {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		log.Panic("ENV err: [" + k + "]" + err.Error())
	}
	return int32(i)
}

// MustGetInt64 will return the env as int64 or panic if it is not present
func MustGetInt64(k string) int64 {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Panic("ENV err: [" + k + "]" + err.Error())
	}
	return i
}

// GetOrFile attempts to resolve 'key' as an environment variable.
// Failing that, it will check to see if '<key>_FILE' exists.
// If so, it will attempt to read from the referenced file to populate a value.
func GetOrFile(envVar string) (string, error) {
	envVarValue := os.Getenv(envVar)
	if envVarValue != "" {
		return envVarValue, nil
	}

	fileVar := envVar + "_FILE"
	fileVarValue := os.Getenv(fileVar)
	if fileVarValue == "" {
		return envVarValue, nil
	}

	fileContents, err := os.ReadFile(fileVarValue)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(fileContents), "\n"), nil
}

// MustGetOrFile attempts to resolve 'key' as an environment variable.
// Failing that, it will check to see if '<key>_FILE' exists.
// If so, it will attempt to read from the referenced file to populate a value.
// If that fails, it will panic.
func MustGetOrFile(envVar string) string {
	v, err := GetOrFile(envVar)
	if err != nil {
		log.Panic("ENV missing, key: " + envVar)
	}
	return v
}

// GetEnvironment returns the value of the 'ENVIRONMENT' environment variable.
// If it is not set, we will default to 'local'
func GetEnvironment() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return "local"
	}
	return env
}
