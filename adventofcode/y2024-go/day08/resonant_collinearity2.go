/*
Author: Marco Lackovic
Date: 2024-12-14

https://adventofcode.com/2024/day/08

--- Day 8, part 2: Resonant Collinearity ---

Watching over your shoulder as you work, one of The Historians asks if you took the effects of resonant harmonics into your calculations.

Whoops!

After updating your model, it turns out that an antinode occurs at any grid position exactly in line with at least two antennas of the same frequency, regardless of distance. This means that some of the new antinodes will occur at the position of each antenna (unless that antenna is the only one of its frequency).

So, these three T-frequency antennas now create many antinodes:

T....#....
...T......
.T....#...
.........#
..#.......
..........
...#......
..........
....#.....
..........

In fact, the three T-frequency antennas are all exactly in line with two antennas, so they are all also antinodes! This brings the total number of antinodes in the above example to 9.

The original example now has 34 antinodes, including the antinodes that appear on every antenna:

##....#....#
.#.#....0...
..#.#0....#.
..##...0....
....0....#..
.#...#A....#
...#..#.....
#....#.#....
..#.....A...
....#....A..
.#........#.
...#......##

Calculate the impact of the signal using this updated model. How many unique locations within the bounds of the map contain an antinode?

*/

package day08

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
	inputMap := make([][]rune, len(lines))
	for i, line := range lines {
		inputMap[i] = []rune(line)
	}
	// clone inputMap to store the antinodes as bits
	antinodes := make([][]bool, len(inputMap))
	for i := 0; i < len(inputMap); i++ {
		antinodes[i] = make([]bool, len(inputMap[i]))
	}
	// iterate over the matrix and check for each pair of antennas if they are in line
	for i := 0; i < len(inputMap); i++ {
		for j := 0; j < len(inputMap[i]); j++ {
			if inputMap[i][j] == '.' {
				continue
			}
			antenna := inputMap[i][j]
			// search for another antenna with the same frequency
			for k := 0; k < len(inputMap); k++ {
				for l := 0; l < len(inputMap[k]); l++ {
					if inputMap[k][l] == '.' {
						continue
					}
					if i == k && j == l {
						continue
					}
					if antenna == inputMap[k][l] {
						prevI, prevJ := i, j
						currI, currJ := k, l
						for {
							antinodes[currI][currJ] = true
							nextI := prevI + (currI-prevI)*2
							nextJ := prevJ + (currJ-prevJ)*2
							if nextI >= 0 && nextI < len(inputMap) && nextJ >= 0 && nextJ < len(inputMap[nextI]) {
								prevI, prevJ = currI, currJ
								currI, currJ = nextI, nextJ
							} else {
								break
							}
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(antinodes); i++ {
		for j := 0; j < len(antinodes[i]); j++ {
			if antinodes[i][j] {
				total++
			}
		}
	}
	return total, nil
}
