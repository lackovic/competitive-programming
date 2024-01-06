package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMirage1GivenExampleShouldReturn114(t *testing.T) {
	// given
	inputFile := "mirage.example"

	// when
	result, err := Solve1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 114, result)
}

func TestMirage2GivenExampleShouldReturn2(t *testing.T) {
	// given
	inputFile := "mirage.example"

	// when
	result, err := Solve2(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 2, result)
}
