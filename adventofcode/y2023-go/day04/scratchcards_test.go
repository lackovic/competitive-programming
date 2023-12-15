package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScratchcards1ShouldReturn13(t *testing.T) {
	// given
	inputFile := "scratchcards.example"

	// when
	result := Solve1(inputFile)

	// then
	assert.Equal(t, 13, result)
}

func TestScratchcards2ShouldReturn30(t *testing.T) {
	// given
	inputFile := "scratchcards.example"

	// when
	result := Solve2(inputFile)

	// then
	assert.Equal(t, 30, result)
}
