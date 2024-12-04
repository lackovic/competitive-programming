package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2WithExample(t *testing.T) {
	// given
	inputFile := "example"

	// when
	result, err := Part2(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 31, result)
}
