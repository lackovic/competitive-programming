/*
Author: Marco Lackovic
Date: 2023-12-20

https://adventofcode.com/2023/day/7

--- Part Two ---
To make things a little more interesting, the Elf introduces one additional rule. Now, J cards are jokers - wildcards that can act like whatever card would make the hand the strongest type possible.

To balance this, J cards are now the weakest individual cards, weaker even than 2. The other cards stay in the same order: A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J.

J cards can pretend to be whatever card is best for the purpose of determining hand type; for example, QJJQ2 is now considered four of a kind. However, for the purpose of breaking ties between two hands of the same type, J is always treated as J, not the card it's pretending to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.

Now, the above example goes very differently:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
KK677 is now the only two pair, making it the second-weakest hand.
T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.
With the new joker rule, the total winnings in this example are 5905.

Using the new joker rule, find the rank of every hand in your set. What are the new total winnings?

*/

package day07

import (
	"adventofcode2023/utils"
	"fmt"
	"sort"
)

func Solve2(filename string) (int64, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	// read the lines and save them in a slice of hands
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		hands[i], err = parseHand(line, determineHandType2)
	}
	if err != nil {
		return 0, err
	}
	cardStrength['J'] = 0
	sort.Sort(ByStrength(hands))
	result := int64(0)
	for i := int64(0); i < int64(len(hands)); i++ {
		result += int64(hands[i].Bid) * (i + 1)
	}
	return result, nil
}

// determineHandType determines the type of a hand considering jacks as wildcards.
func determineHandType2(cards []rune) (HandType, error) {
	counts := make(map[rune]int)
	jCount := 0
	for _, card := range cards {
		if card == 'J' {
			jCount++
			continue
		}
		counts[card]++
	}
	if jCount == 5 {
		return FiveOfAKind, nil
	}
	switch len(counts) {
	case 5:
		return HighCard, nil
	case 4:
		return OnePair, nil
	case 3:
		for _, count := range counts {
			if count == 3 {
				return ThreeOfAKind, nil
			}
		}
		if jCount == 1 || jCount == 2 {
			return ThreeOfAKind, nil
		}
		return TwoPair, nil
	case 2:
		if jCount == 1 {
			for _, count := range counts {
				if count == 3 {
					return FourOfAKind, nil
				}
			}
			return FullHouse, nil
		}
		for _, count := range counts {
			if count == 3 {
				return FullHouse, nil
			}
		}
		return FourOfAKind, nil
	case 1:
		return FiveOfAKind, nil
	}
	return 0, fmt.Errorf("invalid hand: %v", cards)
}
