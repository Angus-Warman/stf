package stf

import (
	"fmt"
	"os"
	"strconv"
)

func EnvString(key string) (string, error) {
	value, ok := os.LookupEnv(key)

	if !ok {
		return "", fmt.Errorf("env var %s not set", key)
	}

	return value, nil
}

func EnvStringOrDefault(key, defaultVal string) string {
	value, ok := os.LookupEnv(key)

	if ok {
		return value
	}

	return defaultVal
}

func EnvStringOrPanic(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		panic(fmt.Sprintf("required env var %s is not set", key))
	}

	return value
}

func EnvInt(key string) (int, error) {
	value, ok := os.LookupEnv(key)

	if !ok {
		return 0, fmt.Errorf("env var %s not set", key)
	}

	i, err := strconv.Atoi(value)

	if err != nil {
		return 0, fmt.Errorf("env var %s %q could not be parsed", key, value)
	}

	return i, nil
}

func EnvIntOrDefault(key string, defaultVal int) int {
	i, err := EnvInt(key)

	if err != nil {
		return defaultVal
	}

	return i
}

func EnvIntOrPanic(key string) int {
	i, err := EnvInt(key)

	if err != nil {
		panic(err)
	}

	return i
}

func EnvBool(key string) bool {
	value, ok := os.LookupEnv(key)

	if !ok {
		return false
	}

	b, err := strconv.ParseBool(value)

	if err != nil {
		return false
	}

	return b
}

// EnvBoolChecked returns (_, false) if the env var is not set,
// or cannot be parsed.
func EnvBoolChecked(key string) (bool, bool) {
	value, ok := os.LookupEnv(key)

	if !ok {
		return false, false
	}

	b, err := strconv.ParseBool(value)

	if err != nil {
		return false, false
	}

	return b, true
}
