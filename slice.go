package stf

import (
	"errors"
)

func Convert[I any, O any](input []I, converter func(I) O) []O {
	output := make([]O, len(input))

	for i, inputValue := range input {
		output[i] = converter(inputValue)
	}

	return output
}

func ConvertE[I any, O any](input []I, converter func(I) (O, error)) ([]O, error) {
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

func First[T any](input []T, isPermitted func(T) bool) (T, bool) {
	var zero T

	for _, value := range input {
		if isPermitted(value) {
			return value, true
		}
	}

	return zero, false
}

func All[T any](input []T, isPermitted func(T) bool) bool {
	for _, value := range input {
		if !isPermitted(value) {
			return false
		}
	}

	return true
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

func Chunk[T any](slice []T, numChunks int) [][]T {
	if numChunks < 1 {
		return [][]T{}
	}

	chunks := make([][]T, numChunks)
	avgChunkSize := len(slice) / numChunks    // Most chunks will be this size
	numLargerChunks := len(slice) % numChunks // This many chunks will be one larger

	start := 0
	for i := range numChunks {
		end := start + avgChunkSize
		if i < numLargerChunks {
			end++
		}
		chunks[i] = slice[start:end]
		start = end
	}

	return chunks
}
