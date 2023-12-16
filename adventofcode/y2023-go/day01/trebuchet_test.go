package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPassingTheExampleInputShouldReturn142(t *testing.T) {
	// given
	inputFile := "trebuchet.example"

	// when
	result, err := Solve(inputFile, 1)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 142, result)
}

func TestWhenPassingThePart2ExampleInputShouldReturn281(t *testing.T) {
	// given
	inputFile := "trebuchet_part2.example"

	// when
	result, err := Solve(inputFile, 2)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 281, result)
}
