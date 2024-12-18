package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IntToArrayOfBits generates a slice of bits (0s and 1s) based on the given number and length.
func IntToArrayOfBits(i, length int) []int {
	bits := make([]int, length)
	for j := 0; j < length; j++ {
		bits[j] = (i >> uint(j)) & 1
	}
	return bits
}

// IntToArrayOfBits generates a slice of (0s, 1s and 2s) based on the given number and length.
func IntToArrayOfTrits(i, length int) []int {
	trits := make([]int, length)
	for j := 0; j < length; j++ {
		trits[j] = i % 3
		i /= 3
	}
	return trits
}

// IntToSliceOfDigits converts an integer to a slice of its digits.
func IntToSliceOfDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	var digits []int
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}
	return digits
}

// SliceOfDigitsToInt converts a slice of integers to an integer.
func SliceOfDigitsToInt(digits []int) int {
	result, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(digits)), ""), "[]"))
	return result
}
