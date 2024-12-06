/*
Author: Marco Lackovic
Date: 2024-12-05

https://adventofcode.com/2024/day/03

--- Day 3, part 1: Mull It Over ---

"Our computers are having issues, so I have no idea if we have any Chief Historians in stock! You're welcome to check the warehouse, though," says the mildly flustered shopkeeper at the North Pole Toboggan Rental Shop. The Historians head out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are having issues again?"

The computer appears to be trying to run a program, but its memory (your puzzle input) is corrupted. All of the instructions have been jumbled up!

It seems like the goal of the program is just to multiply some numbers. It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also many invalid characters that should be ignored, even if they look like part of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).

Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add up all of the results of the multiplications?

*/

package day03

import (
	"adventofcode2024/utils"
	"fmt"
	"regexp"
)

func Part1(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	total := 0
	for _, line := range lines {
		multiplications, err := multiply(line)
		if err != nil {
			return 0, fmt.Errorf("failed to multiply: %w", err)
		}
		total += multiplications
	}
	return total, nil
}

func multiply(line string) (int, error) {
	regex, err := regexp.Compile(`mul\(\d+,\d+\)`)
	if err != nil {
		return 0, err
	}
	multiplications := 0
	for i := 0; i < len(line); i++ {
		if line[i] == 'm' {
			for j := i + 1; j < len(line); j++ {
				if line[j] == ')' {
					match := regex.FindString(line[i : j+1])
					if match != "" {
						var a, b int
						_, err := fmt.Sscanf(match, "mul(%d,%d)", &a, &b)
						if err != nil {
							return 0, err
						}
						multiplications += a * b
					}
					i = j
					break
				}
			}
		}
	}
	return multiplications, nil
}
