/*
Author: Marco Lackovic
Date: 2024-12-11

https://adventofcode.com/2024/day/07

--- Day 7, part 2: Bridge Repair ---

The engineers seem concerned; the total calibration result you gave them is nowhere close to being within safety tolerances. Just then, you spot your mistake: some well-hidden elephants are holding a third type of operator.

The concatenation operator (||) combines the digits from its left and right inputs into a single number. For example, 12 || 345 would become 12345. All operators are still evaluated left-to-right.

Now, apart from the three equations that could be made true using only addition and multiplication, the above example has three more equations that can be made true by inserting operators:

- 156: 15 6 can be made true through a single concatenation: 15 || 6 = 156.
- 7290: 6 8 6 15 can be made true using 6 * 8 || 6 * 15.
- 192: 17 8 14 can be made true using 17 || 8 + 14.

Adding up all six test values (the three that could be made before using only + and * plus the new three that can now be made by also using ||) produces the new total calibration result of 11387.

Using your new knowledge of elephant hiding spots, determine which equations could possibly be true. What is their total calibration result?

*/

package day07

import (
	"adventofcode2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part2(filename string) (int, error) {
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
		result := findMatchingExpression(numbers, testValue)
		total += result
	}
	return total, nil
}

// findMatchingExpression uses a binary representation of len(numbers) to cycle through all possible combinations of operators, using 0 for +, 1 for * and 2 for ||.
func findMatchingExpression(numbers []int64, testValue int64) int {
	if len(numbers) == 1 {
		if testValue == numbers[0] {
			return int(testValue)
		}
		return 0
	}
	limit := 1
	for i := 0; i < len(numbers)-1; i++ {
		limit *= 3
	}
	for i := 0; i < limit; i++ {
		result := calculateExpression(i, numbers)
		if result == int(testValue) {
			return int(testValue)
		}
	}
	return 0
}

// calculateExpression calculates the result of the expression based on the given operators, where 0 is +, 1 is * and 2 is ||
func calculateExpression(i int, numbers []int64) int {
	operators := utils.IntToArrayOfTrits(i, len(numbers)-1)
	result := int(numbers[0])
	for j := 0; j < len(operators); j++ {
		if operators[j] == 0 {
			result += int(numbers[j+1])
		} else if operators[j] == 1 {
			result *= int(numbers[j+1])
		} else {
			newNumber, _ := utils.StringToInt64(strconv.Itoa(int(result)) + strconv.Itoa(int(numbers[j+1])))
			result = int(newNumber)
		}
	}
	return result
}
