package day12

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
	assert.Equal(t, 1206, result)
}

func TestPart2WithExampleE(t *testing.T) {
	// given
	inputFile := "exampleE"

	// when
	result, err := Part2(inputFile)
	if err != nil {
		t.Error(err)
		return
	}

	// then
	assert.Equal(t, 236, result)
}
