package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPassingTheExampleInputShouldReturn8(t *testing.T) {
	// given
	inputFile := "cube_conundrum.example"

	// when
	result := Solve(inputFile)

	// then
	assert.Equal(t, 8, result)
}

func TestWhenPassingTheExampleInputShouldReturn2286(t *testing.T) {
	// given
	inputFile := "cube_conundrum.example"

	// when
	result := SolvePart2(inputFile)

	// then
	assert.Equal(t, 2286, result)
}
