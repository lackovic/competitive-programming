package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {
	// given
	inputFile := "example1"

	// when
	result, err := Part1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 161, result)
}
