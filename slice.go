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
