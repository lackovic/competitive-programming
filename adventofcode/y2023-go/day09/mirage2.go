/*
Author: Marco Lackovic Date: 2024-01-06

https://adventofcode.com/2023/day/9

--- Part Two ---

Of course, it would be nice to have even more history included in your report.
Surely it's safe to just extrapolate backwards as well, right?

For each history, repeat the process of finding differences until the sequence
of differences is entirely zero. Then, rather than adding a zero to the end and
filling in the next values of each previous sequence, you should instead add a
zero to the beginning of your sequence of zeroes, then fill in new first values
for each previous sequence.

In particular, here is what the third example history looks like when
extrapolating back in time:

5  10  13  16  21  30  45
  5   3   3   5   9  15
   -2   0   2   4   6
      2   2   2   2
        0   0   0

Adding the new values on the left side of each sequence from bottom to top
eventually reveals the new left-most history value: 5.

Doing this for the remaining example data above results in previous values of -3
for the first history and 0 for the second history. Adding all three new values
together produces 2.

Analyze your OASIS report again, this time extrapolating the previous value for
each history. What is the sum of these extrapolated values?

*/

package day09

import (
	"adventofcode2023/utils"
	"errors"
)

func Solve2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, history := range lines {
		p, err := extrapolateBackwards(history)
		if err != nil {
			return 0, err
		}
		sum += p
	}
	return sum, nil
}

// extrapolateBackwards returns the extrapolated previous value in the history.
func extrapolateBackwards(history string) (int, error) {
	values, err := utils.StringToSliceOfInts(history)
	if err != nil {
		return 0, err
	}
	sequences := [][]int{values}
	for i := 0; i < len(values); i++ {
		diff := sequences[i][1] - sequences[i][0]
		sum := diff
		sequences = append(sequences, []int{diff})
		for j := 1; j < len(sequences[i])-1; j++ {
			diff = sequences[i][j+1] - sequences[i][j]
			sequences[i+1] = append(sequences[i+1], diff)
			sum += diff
		}
		if sum == 0 {
			extrapolation := 0
			for j := i; j >= 0; j-- {
				extrapolation = sequences[j][0] - extrapolation
			}
			return extrapolation, nil
		}
	}
	return 0, errors.New("extrapolation not found")
}
