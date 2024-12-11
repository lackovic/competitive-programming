/*
Author: Marco Lackovic
Date: 2024-12-09

https://adventofcode.com/2024/day/07

--- Day 7, part 1: Bridge Repair ---

The Historians take you to a familiar rope bridge over a river in the middle of a jungle. The Chief isn't on this side of the bridge, though; maybe he's on the other side?

When you go to cross the bridge, you notice a group of engineers trying to repair it. (Apparently, it breaks pretty frequently.) You won't be able to cross until it's fixed.

You ask how long it'll take; the engineers tell you that it only needs final calibrations, but some young elephants were playing nearby and stole all the operators from their calibration equations! They could finish the calibrations if only someone could determine which test values could possibly be produced by placing any combination of operators into their calibration equations (your puzzle input).

For example:

190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20

Each line represents a single equation. The test value appears before the colon on each line; it is your job to determine whether the remaining numbers can be combined with operators to produce the test value.

Operators are always evaluated left-to-right, not according to precedence rules. Furthermore, numbers in the equations cannot be rearranged. Glancing into the jungle, you can see elephants holding two different types of operators: add (+) and multiply (*).

Only three of the above equations can be made true by inserting operators:

- 190: 10 19 has only one position that accepts an operator: between 10 and 19. Choosing + would give 29, but choosing * would give the test value (10 * 19 = 190).
- 3267: 81 40 27 has two positions for operators. Of the four possible configurations of the operators, two cause the right side to match the test value: 81 + 40 * 27 and 81 * 40 + 27 both equal 3267 (when evaluated left-to-right)!
- 292: 11 6 16 20 can be solved in exactly one way: 11 + 6 * 16 + 20.

The engineers just need the total calibration result, which is the sum of the test values from just the equations that could possibly be true. In the above example, the sum of the test values for the three equations listed above is 3749.

Determine which equations could possibly be true. What is their total calibration result?

*/

package day07

import (
	"adventofcode2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	total := 0
	for _, line := range lines {
		equation := strings.Split(line, ":")
		testValue, err := strconv.ParseInt(strings.TrimSpace(equation[0]), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to convert test value to int: %w", err)
		}
		numbers, err := utils.StringToSliceOfInts64(strings.TrimSpace(equation[1]))
		if err != nil {
			return 0, fmt.Errorf("failed to convert numbers to slice of ints: %w", err)
		}
		if len(numbers) == 1 {
			if testValue == numbers[0] {
				total += int(testValue)
				continue
			}
		}
		// use a binary representation of len(numbers) to cycle through all possible combinations of operators, where 0 is + and 1 is *
		for i := 0; i < 1<<uint(len(numbers)-1); i++ {
			operators := utils.IntToArrayOfBits(i, len(numbers)-1)
			result := int(numbers[0])
			for j := 0; j < len(operators); j++ {
				if operators[j] == 0 {
					result += int(numbers[j+1])
				} else {
					result *= int(numbers[j+1])
				}
			}
			if result == int(testValue) {
				total += int(testValue)
				break
			}
		}
	}
	return total, nil
}
