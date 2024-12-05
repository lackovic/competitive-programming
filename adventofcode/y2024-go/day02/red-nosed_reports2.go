/*
Author: Marco Lackovic
Date: 2024-12-05

https://adventofcode.com/2024/day/02

--- Day 2, part 2: Red-Nosed Reports ---

The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the above example's reports are now safe:

- 7 6 4 2 1: Safe without removing any level.
- 1 2 7 8 9: Unsafe regardless of which level is removed.
- 9 7 6 2 1: Unsafe regardless of which level is removed.
- 1 3 2 4 5: Safe by removing the second level, 3.
- 8 6 4 4 1: Safe by removing the third level, 4.
- 1 3 6 7 9: Safe without removing any level.

Thanks to the Problem Dampener, 4 reports are actually safe!

Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?

*/

package day02

import (
	"adventofcode2024/utils"
	"fmt"
	"slices"
)

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	total := 0
	for _, line := range lines {
		report, err := utils.StringToSliceOfInts(line)
		if err != nil {
			return 0, fmt.Errorf("failed to convert string to slice of ints: %w", err)
		}
		safe, err := isSafeWithTolerance(report)
		if err != nil {
			return 0, fmt.Errorf("failed to check if report is safe: %w", err)
		}
		if safe {
			total++
		}
	}
	return total, nil
}

func isSafeWithTolerance(levels []int) (bool, error) {
	if levels[0] == levels[1] {
		levels = levels[1:]
		return isSafe(levels)
	}
	increasing := levels[0] < levels[1]
	for i := 0; i < len(levels)-1; i++ {
		if i == len(levels)-2 {
			return true, nil
		}
		diff := levels[i+1] - levels[i]
		if increasing {
			if diff < 1 || diff > 3 {
				return isSafeNoTolerance(levels, i)
			}
		} else if diff < -3 || diff > -1 {
			return isSafeNoTolerance(levels, i)
		}
	}
	return true, nil
}

func isSafeNoTolerance(levels []int, i int) (bool, error) {
	copy := slices.Clone(levels)
	safe, err := isSafe(append(copy[:i], copy[i+1:]...))
	if err != nil {
		return false, fmt.Errorf("failed to check if report is safe: %w", err)
	}
	if safe {
		return true, nil
	}
	copy = slices.Clone(levels)
	safe, err = isSafe(append(copy[:i+1], copy[i+2:]...))
	if err != nil {
		return false, fmt.Errorf("failed to check if report is safe: %w", err)
	}
	if safe {
		return true, nil
	}
	if i == 0 {
		return false, nil
	}
	return isSafe(append(levels[:i-1], levels[i:]...))
}
