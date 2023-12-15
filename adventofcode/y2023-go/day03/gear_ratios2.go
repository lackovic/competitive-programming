/*
Author: Marco Lackovic
Date: 2023-12-13

https://adventofcode.com/2023/day/3

--- Part Two ---

The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?

*/

package day03

import (
	"adventofcode2023/utils"
	"strconv"
)

func Solve2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	total := findGearsRatios("", lines[0], lines[1])
	row := 1
	for ; row < len(lines)-1; row++ {
		total += findGearsRatios(lines[row-1], lines[row], lines[row+1])
	}
	total += findGearsRatios(lines[row-1], lines[row], "")
	return total, nil
}

func findGearsRatios(prevLine, currentLine, nextLine string) (total int) {
	nextNumber := 0
	total = 0
	for i := 0; i < len(currentLine); {
		nextNumber, i = findNextGearRatio(i, prevLine, currentLine, nextLine)
		total += nextNumber
	}
	return total
}

// findNextGearRatio returns the next gear	adjacent to exactly two part numbers otherwise it returns 0.
func findNextGearRatio(i int, prevLine, currentLine, nextLine string) (nextNumber int, nextI int) {
	for i < len(currentLine) && !isGear(currentLine[i]) {
		i++
	}
	if i == len(currentLine) {
		return 0, i
	}
	numbers := findSurroundingNumbers(i, prevLine, currentLine, nextLine)
	if len(numbers) == 2 {
		return numbers[0] * numbers[1], i + 1
	}
	return 0, i + 1
}

// findSurroundingNumbers returns the numbers surrounding the given index on the current line
func findSurroundingNumbers(i int, prevLine, currentLine, nextLine string) (numbers []int) {
	start := i
	if start > 0 {
		start--
	}
	end := i + 1
	if end < len(currentLine)-1 {
		end++
	}
	if prevLine != "" {
		numbers = append(numbers, findNumbers(prevLine, start, end)...)
	}
	numbers = append(numbers, findNumbers(currentLine, start, end)...)
	if nextLine != "" {
		numbers = append(numbers, findNumbers(nextLine, start, end)...)
	}
	return numbers
}

// findNumbers returns the numbers in the given string.
func findNumbers(line string, start, end int) (numbers []int) {
	for i := start; i < end; {
		number := 0
		number, i = findNextNumber(line, i, end)
		if number != 0 {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

// findNextNumber returns the next number starting from i if found otherwise it returns 0.
func findNextNumber(line string, i, end int) (nextNumber int, nextI int) {
	for i < end && !isDigit(line[i]) {
		i++
	}
	if i == end {
		return 0, i
	}
	for ; i > 0 && isDigit(line[i-1]); i-- {
	}
	start := i
	for i < len(line) && isDigit(line[i]) {
		i++
	}
	nextNumber, _ = strconv.Atoi(line[start:i])
	return nextNumber, i
}

// isGear returns true if the given char is a gear.
func isGear(char byte) bool {
	return char == '*'
}
