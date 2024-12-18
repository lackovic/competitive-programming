/*
Author: Marco Lackovic
Date: 2024-12-16
https://adventofcode.com/2024/day/11
*/

package day11

import (
	"adventofcode2024/utils"
	"fmt"
	"strconv"
)

func Part1(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	stones, err := utils.StringWithSeparatorToSliceOfInts(lines[0], ' ')
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to slice of ints: %w", err)
	}
	for i := 0; i < 25; i++ {
		newStones := make([]int, 0)
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				digits := utils.IntToSliceOfDigits(stone)
				half := len(digits) / 2
				left := utils.SliceOfDigitsToInt(digits[:half])
				right := utils.SliceOfDigitsToInt(digits[half:])
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return len(stones), nil
}
