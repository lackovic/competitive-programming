/*
Author: Marco Lackovic
Date: 2024-12-06

https://adventofcode.com/2024/day/04

--- Day 4, part 1: Ceres Search ---

"Looks like the Chief's not here. Next!" One of The Historians pulls out a device and pushes the only button on it. After a brief flash, you recognize the interior of the Ceres monitoring station!

As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt; she'd like to know if you could help her with her word search (your puzzle input). She only has to find one word: XMAS.

This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words. It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them. Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:


..X...
.SAMX.
.A..A.
XMAS.S
.X....

The actual word search will be full of letters instead. For example:

MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX

In this word search, XMAS occurs a total of 18 times; here's the same word search again, but where letters not involved in any XMAS have been replaced with .:

....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX

Take a look at the little Elf's word search. How many times does XMAS appear?

*/

package day04

import (
	"adventofcode2024/utils"
	"fmt"
)

func Part1(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	letterMatrix := utils.StringsToChars(lines)
	total := 0
	for i := 0; i < len(letterMatrix); i++ {
		for j := 0; j < len(letterMatrix[i]); j++ {
			if letterMatrix[i][j] == 'X' {
				total += xmasCountInAllDirections(letterMatrix, i, j)
			}
		}
	}
	return total, nil
}

func xmasCountInAllDirections(letterMatrix [][]rune, i, j int) int {
	count := 0
	if i+3 < len(letterMatrix) {
		count += xmasCount(letterMatrix[i][j], letterMatrix[i+1][j], letterMatrix[i+2][j], letterMatrix[i+3][j])
		if j+3 < len(letterMatrix[i]) {
			count += xmasCount(letterMatrix[i][j], letterMatrix[i+1][j+1], letterMatrix[i+2][j+2], letterMatrix[i+3][j+3])
		}
		if j-3 >= 0 {
			count += xmasCount(letterMatrix[i][j], letterMatrix[i+1][j-1], letterMatrix[i+2][j-2], letterMatrix[i+3][j-3])
		}
	}
	if i-3 >= 0 {
		count += xmasCount(letterMatrix[i][j], letterMatrix[i-1][j], letterMatrix[i-2][j], letterMatrix[i-3][j])
		if j+3 < len(letterMatrix[i]) {
			count += xmasCount(letterMatrix[i][j], letterMatrix[i-1][j+1], letterMatrix[i-2][j+2], letterMatrix[i-3][j+3])
		}
		if j-3 >= 0 {
			count += xmasCount(letterMatrix[i][j], letterMatrix[i-1][j-1], letterMatrix[i-2][j-2], letterMatrix[i-3][j-3])
		}
	}
	if j+3 < len(letterMatrix[i]) {
		count += xmasCount(letterMatrix[i][j], letterMatrix[i][j+1], letterMatrix[i][j+2], letterMatrix[i][j+3])
	}
	if j-3 >= 0 {
		count += xmasCount(letterMatrix[i][j], letterMatrix[i][j-1], letterMatrix[i][j-2], letterMatrix[i][j-3])
	}
	return count
}

// given 4 letters, returns 1 if they spell XMAS, 0 otherwise
func xmasCount(a, b, c, d rune) int {
	if a == 'X' && b == 'M' && c == 'A' && d == 'S' {
		return 1
	}
	return 0
}
