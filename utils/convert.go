package utils

import (
	"strconv"
)

func ToIntSlice(input []string) ([]int, error) {
	ints := make([]int, len(input))
	for i, s := range input {
		converted, err := strconv.Atoi(s)

		if err != nil {
			return nil, err
		}

		ints[i] = converted
	}
	return ints, nil
}
