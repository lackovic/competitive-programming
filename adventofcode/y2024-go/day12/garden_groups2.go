/*
Author: Marco Lackovic
Date: 2024-12-18
https://adventofcode.com/2024/day/12
*/

package day12

import (
	"adventofcode2024/utils"
	"fmt"
)

// map of stones to blinking times to stones count
var stonesCountCache = make(map[int]map[int]int)

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	region := utils.StringsToChars(lines)
	visited := make([][]bool, len(region))
	for i := 0; i < len(region); i++ {
		visited[i] = make([]bool, len(region[i]))
	}
	total := 0
	for i := 0; i < len(region); i++ {
		for j := 0; j < len(region[i]); j++ {
			if !visited[i][j] {
				area, cornersCount := getAreaAndCornersCount(region, visited, i, j)
				total += area * cornersCount
			}
		}
	}
	return total, nil
}

// getAreaAndCornersCount returns the area and corners of the garden group that contains the garden plot at (i, j)
func getAreaAndCornersCount(region [][]rune, visited [][]bool, i, j int) (int, int) {
	visited[i][j] = true
	area := 1
	cornersCount := 0
	if i > 0 && region[i-1][j] == region[i][j] {
		if !visited[i-1][j] {
			a, c := getAreaAndCornersCount(region, visited, i-1, j)
			area += a
			cornersCount += c
		}
		if j < len(region[i])-1 && region[i][j+1] == region[i][j] && region[i-1][j+1] != region[i][j] {
			cornersCount++
		}
	} else if j == len(region[i])-1 || region[i][j+1] != region[i][j] {
		cornersCount++
	}
	if i < len(region)-1 && region[i+1][j] == region[i][j] {
		if !visited[i+1][j] {
			a, c := getAreaAndCornersCount(region, visited, i+1, j)
			area += a
			cornersCount += c
		}
		if j > 0 && region[i][j-1] == region[i][j] && region[i+1][j-1] != region[i][j] {
			cornersCount++
		}
	} else if j == 0 || region[i][j-1] != region[i][j] {
		cornersCount++
	}
	if j > 0 && region[i][j-1] == region[i][j] {
		if !visited[i][j-1] {
			a, c := getAreaAndCornersCount(region, visited, i, j-1)
			area += a
			cornersCount += c
		}
		if i > 0 && region[i-1][j] == region[i][j] && region[i-1][j-1] != region[i][j] {
			cornersCount++
		}
	} else if i == 0 || region[i-1][j] != region[i][j] {
		cornersCount++
	}
	if j < len(region[i])-1 && region[i][j+1] == region[i][j] {
		if !visited[i][j+1] {
			a, c := getAreaAndCornersCount(region, visited, i, j+1)
			area += a
			cornersCount += c
		}
		if i < len(region)-1 && region[i+1][j] == region[i][j] && region[i+1][j+1] != region[i][j] {
			cornersCount++
		}
	} else if i == len(region)-1 || region[i+1][j] != region[i][j] {
		cornersCount++
	}
	return area, cornersCount
}
