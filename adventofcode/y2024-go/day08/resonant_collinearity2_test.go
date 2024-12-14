package day08

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
		return
	}

	// then
	assert.Equal(t, 34, result)
}
