package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPassingTheExampleInputShouldReturn142(t *testing.T) {
	// given
	inputFile := "trebuchet.example"

	// when
	result := Solve(inputFile)

	// then
	assert.Equal(t, 142, result)
}
