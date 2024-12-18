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

func Part1(filename string) (int, error) {
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
				area, perimeter := getAreaAndPerimeter(region, visited, i, j)
				total += area * perimeter
			}
		}
	}
	return total, nil
}

// getAreaAndPerimeter returns the area and perimeter of the garden group that contains the garden plot at (i, j)
func getAreaAndPerimeter(region [][]rune, visited [][]bool, i, j int) (int, int) {
	visited[i][j] = true
	area := 1
	perimeter := 0
	if i > 0 && region[i-1][j] == region[i][j] {
		if !visited[i-1][j] {
			a, p := getAreaAndPerimeter(region, visited, i-1, j)
			area += a
			perimeter += p
		}
	} else {
		perimeter++
	}
	if i < len(region)-1 && region[i+1][j] == region[i][j] {
		if !visited[i+1][j] {
			a, p := getAreaAndPerimeter(region, visited, i+1, j)
			area += a
			perimeter += p
		}
	} else {
		perimeter++
	}
	if j > 0 && region[i][j-1] == region[i][j] {
		if !visited[i][j-1] {
			a, p := getAreaAndPerimeter(region, visited, i, j-1)
			area += a
			perimeter += p
		}
	} else {
		perimeter++
	}
	if j < len(region[i])-1 && region[i][j+1] == region[i][j] {
		if !visited[i][j+1] {
			a, p := getAreaAndPerimeter(region, visited, i, j+1)
			area += a
			perimeter += p
		}
	} else {
		perimeter++
	}
	return area, perimeter
}
