package stf

import (
	"fmt"
	"os"
	"strconv"
)

func EnvString(key, defaultVal string) string {
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

func EnvInt(key string, defaultVal int) int {
	raw, ok := os.LookupEnv(key)

	if !ok {
		return defaultVal
	}

	value, err := strconv.Atoi(raw)

	if err != nil {
		return defaultVal
	}

	return value
}

func EnvIntOrPanic(key string) int {
	raw, ok := os.LookupEnv(key)

	if !ok {
		panic(fmt.Errorf("env var %q not set", key))
	}

	value, err := strconv.Atoi(raw)

	if err != nil {
		panic(fmt.Errorf("could not parse %q as int", raw))
	}

	return value
}

func EnvFloat(key string, defaultVal float64) float64 {
	raw, ok := os.LookupEnv(key)

	if !ok {
		return defaultVal
	}

	value, err := strconv.ParseFloat(raw, 64)

	if err != nil {
		return defaultVal
	}

	return value
}

func EnvFloatOrPanic(key string) float64 {
	raw, ok := os.LookupEnv(key)

	if !ok {
		panic(fmt.Errorf("env var %q not set", key))
	}

	value, err := strconv.ParseFloat(raw, 64)

	if err != nil {
		panic(fmt.Errorf("could not parse %q as float", raw))
	}

	return value
}

// Default false
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
	raw, ok := os.LookupEnv(key)

	if !ok {
		return false, false
	}

	value, err := strconv.ParseBool(raw)

	if err != nil {
		return false, false
	}

	return value, true
}
