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
