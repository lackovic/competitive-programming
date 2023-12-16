/*
Author: Marco Lackovic
Date: 2023-12-16

https://adventofcode.com/2023/day/5

--- Part Two ---

Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds: line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and the second value is the length of the range. So, in the first line of the example above:

seeds: 79 14 55 13

This line describes two ranges of seed numbers to be planted in the garden. The first range starts with seed number 79 and contains 14 values: 79, 80, ..., 91, 92. The second range starts with seed number 55 and contains 13 values: 55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first line of the almanac. What is the lowest location number that corresponds to any of the initial seed numbers?

*/

package day05

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
)

func Solve2(filename string) (int64, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	values, err := extractNumbersToInt64(strings.Split(lines[0], ":")[1])
	if err != nil {
		return 0, err
	}
	mapsList, err := extractMaps(lines)
	if err != nil {
		return 0, err
	}
	minValue := int64(1<<63 - 1)
	for j := 0; j < len(values); j += 2 {
		fmt.Printf("seeds group %d of %d\n", j/2+1, len(values)/2)
		currentValue := make([]int64, 1)
		for length := int64(0); length < values[j+1]; length++ {
			currentValue[0] = values[j] + length
			for i := 0; i < len(mapsList); i++ {
				convert(currentValue, mapsList[i])
			}
			if currentValue[0] < minValue {
				minValue = currentValue[0]
				fmt.Println("minValue:", minValue)
			}
		}
	}
	return minValue, nil
}

// extractMaps returns a list of maps, where each map is an array of arrays of int64.
func extractMaps(lines []string) ([][][]int64, error) {
	var mapsList [][][]int64
	var maps [][]int64
	for lineIndex := 2; lineIndex < len(lines); lineIndex++ {
		if strings.Contains(lines[lineIndex], ":") {
			mapsList = append(mapsList, maps)
			maps = nil
		} else {
			converted, err := extractNumbersToInt64(lines[lineIndex])
			if err != nil {
				return nil, err
			}
			maps = append(maps, converted)
		}
	}
	mapsList = append(mapsList, maps)
	return mapsList, nil
}
