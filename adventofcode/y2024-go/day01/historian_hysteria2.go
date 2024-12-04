/*
Author: Marco Lackovic
Date: 2024-12-04

https://adventofcode.com/2024/day/01

--- Day 1, part 2: Historian Hysteria ---

Your analysis only confirmed what everyone feared: the two lists of location IDs are indeed very different.

Or are they?

The Historians can't agree on which group made the mistakes or how to read most of the Chief's handwriting, but in the commotion you notice an interesting detail: a lot of location IDs appear in both lists! Maybe the other numbers aren't location IDs at all but rather misinterpreted handwriting.

This time, you'll need to figure out exactly how often each number from the left list appears in the right list. Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.

Here are the same example lists again:

3   4
4   3
2   5
1   3
3   9
3   3

For these example lists, here is the process of finding the similarity score:

- The first number in the left list is 3. It appears in the right list three times, so the similarity score increases by 3 * 3 = 9.
- The second number in the left list is 4. It appears in the right list once, so the similarity score increases by 4 * 1 = 4.
- The third number in the left list is 2. It does not appear in the right list, so the similarity score does not increase (2 * 0 = 0).
- The fourth number, 1, also does not appear in the right list.
- The fifth number, 3, appears in the right list three times; the similarity score increases by 9.
- The last number, 3, appears in the right list three times; the similarity score again increases by 9.

So, for these example lists, the similarity score at the end of this process is 31 (9 + 4 + 0 + 0 + 9 + 9).

Once again consider your left and right lists. What is their similarity score?

*/

package day01

import (
	"adventofcode2024/utils"
	"fmt"
)

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	total := 0
	left, right := []int{}, []int{}
	for _, line := range lines {
		values, err := utils.StringToSliceOfInts(line)
		if err != nil {
			return 0, fmt.Errorf("failed to convert string to slice of ints: %w", err)
		}
		left = append(left, values[0])
		right = append(right, values[1])
	}
	for i := 0; i < len(left); i++ {
		total += left[i] * countNumberInList(left[i], right)
	}
	return total, nil
}

// Counts the number of times a number appears in a given list of numbers
func countNumberInList(number int, list []int) int {
	count := 0
	for _, n := range list {
		if n == number {
			count++
		}
	}
	return count
}
