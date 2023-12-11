package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPassingTheExampleInputShouldReturn142(t *testing.T) {
	// given
	inputFile := "trebuchet.example"

	// when
	result := Solve(inputFile, 1)

	// then
	assert.Equal(t, 142, result)
}

func TestWhenPassingThePart2ExampleInputShouldReturn281(t *testing.T) {
	// given
	inputFile := "trebuchet_part2.example"

	// when
	result := Solve(inputFile, 2)

	// then
	assert.Equal(t, 281, result)
}
