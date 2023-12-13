package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGearRatios1ShouldReturn4361(t *testing.T) {
	// given
	inputFile := "gear_ratios.example"

	// when
	result := Solve1(inputFile)

	// then
	assert.Equal(t, 4361, result)
}

func TestGearRatios2ShouldReturn467835(t *testing.T) {
	// given
	inputFile := "gear_ratios.example"

	// when
	result := Solve2(inputFile)

	// then
	assert.Equal(t, 467835, result)
}
