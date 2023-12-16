package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPassingTheExampleInputShouldReturn8(t *testing.T) {
	// given
	inputFile := "cube_conundrum.example"

	// when
	result, err := Solve(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 8, result)
}

func TestWhenPassingTheExampleInputShouldReturn2286(t *testing.T) {
	// given
	inputFile := "cube_conundrum.example"

	// when
	result, err := SolvePart2(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 2286, result)
}
