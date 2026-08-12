package stf

import "fmt"

func MapKeys[K comparable, V any](target map[K]V) []K {
	keys := make([]K, 0, len(target))

	for k := range target {
		keys = append(keys, k)
	}

	return keys
}

func MapValues[K comparable, V any](target map[K]V) []V {
	vals := make([]V, 0, len(target))

	for _, v := range target {
		vals = append(vals, v)
	}

	return vals
}

func MapKeysAndValues[K comparable, V any](target map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(target))
	vals := make([]V, 0, len(target))

	for k, v := range target {
		keys = append(keys, k)
		vals = append(vals, v)
	}

	return keys, vals
}

func MapGetOrDefault[K comparable, V any](target map[K]V, key K, defaultVal V) V {
	value, ok := target[key]

	if ok {
		return value
	}

	return defaultVal
}

func MapGetKeyFromValue[K comparable, V comparable](target map[K]V, value V) (K, bool) {
	for k, v := range target {
		if v == value {
			return k, true
		}
	}
	var zero K
	return zero, false
}

func MapGetOrErr[K comparable, V any](target map[K]V, key K) (V, error) {
	value, ok := target[key]

	if ok {
		return value, nil
	}

	var zero V

	return zero, fmt.Errorf("key %v not found in map", key)
}
