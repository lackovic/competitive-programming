package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {
	// given
	inputFile := "example"

	// when
	result, err := Part1(inputFile)
	if err != nil {
		t.Error(err)
		return
	}

	// then
	assert.Equal(t, 3749, result)
}
