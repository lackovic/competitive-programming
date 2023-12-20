package day07

import "fmt"

// Hand represents a hand of cards.
type Hand struct {
	Cards []rune
	Type  HandType
	Bid   int
}

func (h Hand) String() string {
	return fmt.Sprintf("%s = %v", string(h.Cards), h.Type)
}

// HandType represents the type of a hand.
type HandType int

func (h HandType) String() string {
	names := [...]string{
		"HighCard",
		"OnePair",
		"TwoPair",
		"ThreeOfAKind",
		"FullHouse",
		"FourOfAKind",
		"FiveOfAKind",
	}
	if h < HighCard || h > FourOfAKind {
		return "Unknown"
	}
	return names[h]
}

// The different types of hands, ordered from weakest to strongest.
const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

// ByStrength implements sort.Interface for []Hand based on the strength of the hands.
type ByStrength []Hand

func (h ByStrength) Len() int {
	return len(h)
}

func (h ByStrength) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (hands ByStrength) Less(i, j int) bool {
	if hands[i].Type != hands[j].Type {
		return hands[i].Type < hands[j].Type
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nerror >>> %s %s <<< error\n\n", hands[i], hands[j])
			panic(r)
		}
	}()
	for k := 0; k < len(hands[i].Cards); k++ {
		if cardStrength[hands[i].Cards[k]] != cardStrength[hands[j].Cards[k]] {
			return cardStrength[hands[i].Cards[k]] < cardStrength[hands[j].Cards[k]]
		}
	}
	return false
}

// cardStrength maps a card to its strength.
var cardStrength = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}
