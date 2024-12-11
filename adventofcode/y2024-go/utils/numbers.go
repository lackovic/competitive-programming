package utils

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
