package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWasteland1GivenExample1ShouldReturn2(t *testing.T) {
	// given
	inputFile := "wasteland.example1"

	// when
	result, err := Solve1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 2, result)
}

func TestWasteland1GivenExample2ShouldReturn6(t *testing.T) {
	// given
	inputFile := "wasteland.example2"

	// when
	result, err := Solve1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 6, result)
}

func TestWasteland2GivenExample1ShouldReturn6(t *testing.T) {
	// given
	inputFile := "wasteland.example3"

	// when
	result, err := Solve2(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, int64(6), result)
}
