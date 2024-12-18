/*
Author: Marco Lackovic
Date: 2024-12-18
https://adventofcode.com/2024/day/11
*/

package day11

import (
	"adventofcode2024/utils"
	"fmt"
	"strconv"
)

// map of stones to blinking times to stones count
var stonesCountCache = make(map[int]map[int]int)

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	stones, err := utils.StringWithSeparatorToSliceOfInts(lines[0], ' ')
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to slice of ints: %w", err)
	}
	stonesCount := 0
	for _, stone := range stones {
		stonesCount += stonesCountAfterBlinking(stone, 75)
	}
	return stonesCount, nil
}

// stonesCountAfterBlinking returns the number of stones after blinking i times.
func stonesCountAfterBlinking(stone, i int) int {
	if stonesCountCache[stone] == nil {
		stonesCountCache[stone] = make(map[int]int)
	}
	// if the number of stones after blinking i times is already calculated, return it
	if stonesCountCache[stone][i] != 0 {
		return stonesCountCache[stone][i]
	}
	if i == 0 {
		stonesCountCache[stone][i] = 1
	} else if stone == 0 {
		stonesCountCache[stone][i] = stonesCountAfterBlinking(1, i-1)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		digits := utils.IntToSliceOfDigits(stone)
		half := len(digits) / 2
		left := utils.SliceOfDigitsToInt(digits[:half])
		right := utils.SliceOfDigitsToInt(digits[half:])
		stonesCountCache[stone][i] = stonesCountAfterBlinking(left, i-1) + stonesCountAfterBlinking(right, i-1)
	} else {
		stonesCountCache[stone][i] = stonesCountAfterBlinking(stone*2024, i-1)
	}
	return stonesCountCache[stone][i]
}
