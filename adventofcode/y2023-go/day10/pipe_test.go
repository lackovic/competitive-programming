package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipe1GivenExample1ShouldReturn4(t *testing.T) {
	// given
	inputFile := "pipe.example1"

	// when
	result, err := Solve1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 4, result)
}

func TestPipe1GivenExample2ShouldReturn8(t *testing.T) {
	// given
	inputFile := "pipe.example2"

	// when
	result, err := Solve1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 8, result)
}
