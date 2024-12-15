/*
Author: Marco Lackovic
Date: 2024-12-15

https://adventofcode.com/2024/day/10

--- Day 10, part 2: Hoof It ---

The reindeer spends a few minutes reviewing your hiking trail map before realizing something, disappearing for a few minutes, and finally returning with yet another slightly-charred piece of paper.

The paper describes a second way to measure a trailhead called its rating. A trailhead's rating is the number of distinct hiking trails which begin at that trailhead. For example:

.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....

The above map has a single trailhead; its rating is 3 because there are exactly three distinct hiking trails which begin at that position:

.....0.   .....0.   .....0.
..4321.   .....1.   .....1.
..5....   .....2.   .....2.
..6....   ..6543.   .....3.
..7....   ..7....   .....4.
..8....   ..8....   ..8765.
..9....   ..9....   ..9....

Here is a map containing a single trailhead with rating 13:

..90..9
...1.98
...2..7
6543456
765.987
876....
987....

This map contains a single trailhead with rating 227 (because there are 121 distinct hiking trails that lead to the 9 on the right edge and 106 that lead to the 9 on the bottom edge):

012345
123456
234567
345678
4.6789
56789.

Here's the larger example from before:

89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732

Considering its trailheads in reading order, they have ratings of 20, 24, 10, 4, 1, 4, 5, 8, and 5. The sum of all trailhead ratings in this larger example topographic map is 81.

You're not sure how, but the reindeer seems to have crafted some tiny flags out of toothpicks and bits of paper and is using them to mark trailheads on your topographic map. What is the sum of the ratings of all trailheads?

*/

package day10

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
	topographicMap := utils.StringsToInts(lines)
	for i := 0; i < len(topographicMap); i++ {
		for j := 0; j < len(topographicMap[i]); j++ {
			if topographicMap[i][j] == 0 {
				total += calculateRating(topographicMap, i, j, 0)
			}
		}
	}
	return total, nil
}

// calculateRating returns the number of trails starting at the given position (i, j) and height start
func calculateRating(topographicMap [][]int, i, j, start int) int {
	if i < 0 || i >= len(topographicMap) || j < 0 || j >= len(topographicMap[i]) {
		return 0
	}
	if topographicMap[i][j] != start {
		return 0
	}
	if topographicMap[i][j] == 9 {
		return 1
	}
	return calculateRating(topographicMap, i+1, j, start+1) + calculateRating(topographicMap, i-1, j, start+1) + calculateRating(topographicMap, i, j+1, start+1) + calculateRating(topographicMap, i, j-1, start+1)
}
