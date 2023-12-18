package utils

import (
	"strconv"
	"strings"
)

// StringToArrayOfInts given a string of numbers separated by spaces, returns a slice of ints.
func StringToSliceOfInts(s string) ([]int, error) {
	fields := strings.Fields(s)
	ints := make([]int, len(fields))
	for i, field := range fields {
		var err error
		ints[i], err = strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}
