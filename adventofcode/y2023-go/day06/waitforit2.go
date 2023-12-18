/*
Author: Marco Lackovic
Date: 2023-12-18

https://adventofcode.com/2023/day/6

--- Part Two ---
As the race is about to start, you realize the piece of paper with race times and record distances you got earlier actually just has very bad kerning. There's really only one race - ignore the spaces between the numbers on each line.

So, the example from before:

Time:      7  15   30
Distance:  9  40  200

...now instead means this:

Time:      71530
Distance:  940200

Now, you have to figure out how many ways there are to win this single race. In this example, the race lasts for 71530 milliseconds and the record distance you need to beat is 940200 millimeters. You could hold the button anywhere from 14 to 71516 milliseconds and beat the record, a total of 71503 ways!

How many ways can you beat the record in this one much longer race?

*/

package day06

import (
	"adventofcode2023/utils"
	"strings"
)

func Solve2(filename string) (int64, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}

	time, err := utils.StringToInt64(strings.Split(lines[0], ":")[1])
	if err != nil {
		return 0, err
	}
	distance, err := utils.StringToInt64(strings.Split(lines[1], ":")[1])
	if err != nil {
		return 0, err
	}
	result := int64(1)
	numberOfWaysToBeatRecord := int64(0)
	for holdButtonMs := int64(1); holdButtonMs < time; holdButtonMs++ {
		if holdButtonMs*(time-holdButtonMs) > distance {
			numberOfWaysToBeatRecord++
		}
	}
	result *= numberOfWaysToBeatRecord
	return result, nil
}
