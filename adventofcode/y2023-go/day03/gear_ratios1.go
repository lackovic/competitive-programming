/*
Author: Marco Lackovic
Date: 2023-12-12

https://adventofcode.com/2023/day/2

--- Day 3: Gear Ratios ---

You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

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

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?

*/

package day03

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Solve1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		return 0
	}
	defer file.Close()
	schematic, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error reading file:", err)
		return 0
	}
	lines := strings.Split(string(schematic), "\n")
	total := partNumbersSum("", lines[0], lines[1])
	row := 1
	for ; row < len(lines)-1; row++ {
		total += partNumbersSum(lines[row-1], lines[row], lines[row+1])
	}
	total += partNumbersSum(lines[row-1], lines[row], "")
	return total
}

func partNumbersSum(prevLine, currentLine, nextLine string) (total int) {
	nextNumber := 0
	for i := 0; i < len(currentLine); {
		nextNumber, i = findNextNumberSurroundedByAtLeastOneSymbol(i, prevLine, currentLine, nextLine)
		total += nextNumber
	}
	return total
}

// findNextNumberSurroundedByAtLeastOneSymbol returns the next number if it is surrounded by at least one symbol otherwise it returns 0.
func findNextNumberSurroundedByAtLeastOneSymbol(i int, prevLine, currentLine, nextLine string) (nextNumber int, nextI int) {
	for i < len(currentLine) && !isDigit(currentLine[i]) {
		i++
	}
	if i == len(currentLine) {
		return 0, i
	}
	start := i
	for i < len(currentLine) && isDigit(currentLine[i]) {
		i++
	}
	nextNumber, _ = strconv.Atoi(currentLine[start:i])
	if start > 0 {
		start--
	}
	end := i - 1
	if end < len(currentLine)-1 {
		end++
	}
	if isSymbol(currentLine[start]) || isSymbol(currentLine[end]) {
		return nextNumber, i
	}
	if prevLine != "" && containsASymbol(prevLine[start:end+1]) {
		return nextNumber, i
	}
	if nextLine != "" && containsASymbol(nextLine[start:end+1]) {
		return nextNumber, i
	}
	return 0, i
}

// isDigit returns true if the given char is a digit.
func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

// isSymbol returns true if the given char is not a digit and not a period.
func isSymbol(char byte) bool {
	return char != '.' && !isDigit(char)
}

// containsASymbol returns true if the given string contains at least one symbol.
func containsASymbol(line string) bool {
	for i := 0; i < len(line); i++ {
		if isSymbol(line[i]) {
			return true
		}
	}
	return false
}
