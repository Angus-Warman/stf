package stf

import (
	"fmt"
	"os"
	"strconv"
)

func EnvGetOrDefault(key, defaultVal string) string {
	value, ok := os.LookupEnv(key)

	if ok {
		return value
	}

	return defaultVal
}

func EnvGetOrPanic(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		panic(fmt.Sprintf("required env var %s is not set", key))
	}

	return value
}

func EnvGetInt(key string, defaultVal int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return defaultVal
	}

	i, err := strconv.Atoi(value)

	if err != nil {
		return defaultVal
	}

	return i
}

func EnvGetIntOrPanic(key string) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		panic(fmt.Sprintf("required env var %s is not set", key))
	}

	i, err := strconv.Atoi(value)

	if err != nil {
		panic(fmt.Sprintf("env var %s is not a valid int: %s", key, value))
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
