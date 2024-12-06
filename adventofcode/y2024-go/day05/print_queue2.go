/*
Author: Marco Lackovic
Date: 2024-12-06

https://adventofcode.com/2024/day/05

--- Day 5, part 2: Print Queue ---

While the Elves get to work printing the correctly-ordered updates, you have a little time to fix the rest of them.

For each of the incorrectly-ordered updates, use the page ordering rules to put the page numbers in the right order. For the above example, here are the three incorrectly-ordered updates and their correct orderings:

- 75,97,47,61,53 becomes 97,75,47,61,53.
- 61,13,29 becomes 61,29,13.
- 97,13,75,29,47 becomes 97,75,47,29,13.

After taking only the incorrectly-ordered updates and ordering them correctly, their middle page numbers are 47, 29, and 47. Adding these together produces 123.

Find the updates which are not in the correct order. What do you get if you add up the middle page numbers after correctly ordering just those updates?

*/

package day05

import (
	"adventofcode2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	total := 0
	for i, line := range lines {
		// if line doesn't contain a pipe, it's the end of the page ordering rules
		if !strings.Contains(line, "|") {
			lines = lines[i:]
			break
		}
		rule := strings.FieldsFunc(line, func(r rune) bool {
			return r == '|'
		})
		ruleLeft, err := strconv.Atoi(rule[0])
		if err != nil {
			return 0, err
		}
		ruleRight, err := strconv.Atoi(rule[1])
		if err != nil {
			return 0, err
		}
		pageOrderingRules[ruleLeft] = append(pageOrderingRules[ruleLeft], ruleRight)
	}
	for _, line := range lines {
		pages, err := utils.StringWithSeparatorToSliceOfInts(line, ',')
		if err != nil {
			return 0, err
		}
		total += middlePageOfCorrectedUpdate(pages)
	}
	return total, nil
}

// returns the middle page number of an incorrect update after having corrected it, otherwise returns 0
func middlePageOfCorrectedUpdate(pages []int) int {
	correctUpdate := true
	for i := 0; i < len(pages)-1; i++ {
		if shouldBeToTheRight(pages[i], pages[i+1:]) {
			correctUpdate = false
			for j := i + 1; j < len(pages); j++ {
				if slices.Contains(pageOrderingRules[pages[j]], pages[i]) {
					pages[i], pages[j] = pages[j], pages[i]
					i = i - 1
					break
				}
			}
		}
	}
	if correctUpdate {
		return 0
	} else {
		return pages[len(pages)/2]
	}
}
