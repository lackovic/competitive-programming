package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFertilizer1ShouldReturn35(t *testing.T) {
	// given
	inputFile := "fertilizer.example"

	// when
	result, err := Solve1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, int64(35), result)
}

func TestFertilizer2ShouldReturn46(t *testing.T) {
	// given
	inputFile := "fertilizer.example"

	// when
	result, err := Solve2(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, int64(46), result)
}
