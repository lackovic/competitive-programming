package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaitforit1ShouldReturn288(t *testing.T) {
	// given
	inputFile := "waitforit.example"

	// when
	result, err := Solve1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, 288, result)
}
