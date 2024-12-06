/*
Author: Marco Lackovic
Date: 2024-12-06

https://adventofcode.com/2024/day/04

--- Day 4, part 2: Ceres Search ---

The Elf looks quizzically at you. Did you misunderstand the assignment?

Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that is like this:

M.S
.A.
M.S

Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written forwards or backwards.

Here's the same example from before, but this time all of the X-MASes have been kept instead:

.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........

In this example, an X-MAS appears 9 times.

Flip the word search from the instructions back over to the word search side and try again. How many times does an X-MAS appear?

*/

package day04

import (
	"adventofcode2024/utils"
	"fmt"
)

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	letterMatrix := utils.StringsToChars(lines)
	total := 0
	for i := 1; i < len(letterMatrix)-1; i++ {
		for j := 1; j < len(letterMatrix[i])-1; j++ {
			if letterMatrix[i][j] == 'A' {
				total += masCrosses(letterMatrix, i, j)
			}
		}
	}
	return total, nil
}

// returns 1 if the letterMatrix contains the word MAS obliquely ascending and descending, with A in position i, j, 0 otherwise
func masCrosses(letterMatrix [][]rune, i, j int) int {
	if letterMatrix[i-1][j-1] == 'M' {
		if letterMatrix[i+1][j+1] != 'S' {
			return 0
		}
	} else if letterMatrix[i-1][j-1] == 'S' {
		if letterMatrix[i+1][j+1] != 'M' {
			return 0
		}
	} else {
		return 0
	}
	if letterMatrix[i-1][j+1] == 'M' {
		if letterMatrix[i+1][j-1] != 'S' {
			return 0
		}
	} else if letterMatrix[i-1][j+1] == 'S' {
		if letterMatrix[i+1][j-1] != 'M' {
			return 0
		}
	} else {
		return 0
	}
	return 1
}
