/*
Author: Marco Lackovic
Date: 2024-12-06

https://adventofcode.com/2024/day/03

--- Day 3, part 2: Mull It Over ---

As you scan through the corrupted memory, you notice that some of the conditional statements are also still intact. If you handle some of the uncorrupted conditional statements in the program, you might be able to get an even more accurate result.

There are two new instructions you'll need to handle:

- The do() instruction enables future mul instructions.
- The don't() instruction disables future mul instructions.

Only the most recent do() or don't() instruction applies. At the beginning of the program, mul instructions are enabled.

For example:

xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))

This corrupted memory is similar to the example from before, but this time the mul(5,5) and mul(11,8) instructions are disabled because there is a don't() instruction before them. The other mul instructions function normally, including the one at the end that gets re-enabled by a do() instruction.

This time, the sum of the results is 48 (2*4 + 8*5).

Handle the new instructions; what do you get if you add up all of the results of just the enabled multiplications?

*/

package day03

import (
	"adventofcode2024/utils"
	"fmt"
	"regexp"
)

var enabled bool = true

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	total := 0
	for _, line := range lines {
		multiplications, err := multiplyConditionally(line)
		if err != nil {
			return 0, err
		}
		total += multiplications
	}
	return total, nil
}

func multiplyConditionally(line string) (int, error) {
	mulRegex, err := regexp.Compile(`^mul\(\d+,\d+\)$`)
	if err != nil {
		return 0, err
	}
	doRegex, err := regexp.Compile(`^do\(\)$`)
	if err != nil {
		return 0, err
	}
	dontRegex, err := regexp.Compile(`^don't\(\)$`)
	if err != nil {
		return 0, err
	}
	multiplications := 0
	for i := 0; i < len(line); i++ {
		if line[i] == 'd' {
			for j := i + 1; j < len(line); j++ {
				if line[j] == ')' {
					if j-i == 3 {
						match := doRegex.FindString(line[i : j+1])
						if match != "" {
							i = j
							enabled = true
						}
					} else if j-i == 6 {
						match := dontRegex.FindString(line[i : j+1])
						if match != "" {
							i = j
							enabled = false
						}
					}
					break
				}
			}
		} else if enabled && line[i] == 'm' {
			for j := i + 1; j < len(line); j++ {
				if line[j] == ')' {
					if mulRegex.MatchString(line[i : j+1]) {
						var a, b int
						_, err := fmt.Sscanf(line[i:j+1], "mul(%d,%d)", &a, &b)
						if err != nil {
							return 0, err
						}
						multiplications += a * b
					}
					break
				}
			}
		}
	}
	return multiplications, nil
}
