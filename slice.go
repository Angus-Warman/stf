package stf

import "errors"

func Map[I any, O any](input []I, converter func(I) O) []O {
	output := make([]O, len(input))

	for i, inputValue := range input {
		output[i] = converter(inputValue)
	}

	return output
}

func MapE[I any, O any](input []I, converter func(I) (O, error)) ([]O, error) {
	output := []O{}
	errs := []error{}

	for _, inputValue := range input {
		outputValue, err := converter(inputValue)

		if err != nil {
			errs = append(errs, err)
		} else {
			output = append(output, outputValue)
		}
	}

	var err error

	if len(errs) > 0 {
		err = errors.Join(errs...)
	}

	return output, err
}

func Filter[T any](input []T, isPermitted func(T) bool) []T {
	output := []T{}

	for _, inputValue := range input {
		if isPermitted(inputValue) {
			output = append(output, inputValue)
		}
	}

	return output
}

func Split[T any](slice []T, inLeft func(T) bool) (left, right []T) {
	left = make([]T, 0, len(slice))
	right = make([]T, 0, len(slice))
	for _, v := range slice {
		if inLeft(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return
}

func GroupBy[T any, K comparable](slice []T, property func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := property(v)
		result[key] = append(result[key], v)
	}
	return result
}
