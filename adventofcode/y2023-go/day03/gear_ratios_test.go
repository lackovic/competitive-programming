package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGearRatiosShouldReturn4361(t *testing.T) {
	// given
	inputFile := "gear_ratios.example"

	// when
	result := Solve(inputFile)

	// then
	assert.Equal(t, 4361, result)
}
