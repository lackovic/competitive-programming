package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelCards1ShouldReturn6440(t *testing.T) {
	// given
	inputFile := "camel_cards.example"

	// when
	result, err := Solve1(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, int64(6440), result)
}
