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

// StringToInt64 given a string containing a sequence of digits possibly separated by spaces, returns the int64 equivalent to the sequence of digits.
func StringToInt64(s string) (int64, error) {
	s = strings.Replace(s, " ", "", -1)
	return strconv.ParseInt(s, 10, 64)
}

// given an array of strings returns an array of arrays of the characters of the strings
func StringsToChars(ss []string) [][]rune {
	chars := make([][]rune, len(ss))
	for i, s := range ss {
		chars[i] = []rune(s)
	}
	return chars
}
