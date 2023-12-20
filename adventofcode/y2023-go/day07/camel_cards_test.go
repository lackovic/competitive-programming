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

func TestCamelCards2ShouldReturn5905(t *testing.T) {
	// given
	inputFile := "camel_cards.example"

	// when
	result, err := Solve2(inputFile)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, int64(5905), result)
}

func TestDetermineHandType2ShouldReturnFourOfAKind(t *testing.T) {
	// given
	cardsRunes := []rune("J3888")

	// when
	result, err := determineHandType2(cardsRunes)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, HandType(5), result)
}

func TestDetermineHandType2ShouldReturnFullHouse(t *testing.T) {
	// given
	cardsRunes := []rune("A22AA")

	// when
	result, err := determineHandType2(cardsRunes)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, HandType(4), result)
}

func TestDetermineHandType2ShouldReturnFiveOfAKind(t *testing.T) {
	// given
	cardsRunes := []rune("JJJJJ")

	// when
	result, err := determineHandType2(cardsRunes)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, HandType(6), result)
}

func TestDetermineHandType2ShouldReturnOnePair(t *testing.T) {
	// given
	cardsRunes := []rune("JT92A")

	// when
	result, err := determineHandType2(cardsRunes)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, HandType(1), result)
}

func TestDetermineHandType2ShouldReturnThreeOfAKind(t *testing.T) {
	// given
	cardsRunes := []rune("J9QJ5")

	// when
	result, err := determineHandType2(cardsRunes)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, HandType(3), result)
}
