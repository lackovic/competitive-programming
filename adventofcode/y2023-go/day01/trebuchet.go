/*
Author: Marco Lackovic
Date: 2023-12-09

https://adventofcode.com/2023/day/1

--- Day 1: Trebuchet?! ---

Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

--- Part Two ---

Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?

*/

package day01

import (
	"adventofcode2023/utils"
	"fmt"
)

func Solve(filename string, part int) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		calibrationVal, err := calibrationValue(line, part)
		if err != nil {
			return 0, fmt.Errorf("error calculating the calibration value: %w", err)
		}
		total += calibrationVal
	}
	return total, nil
}

// calibrationValue returns the calibration value of a line.
func calibrationValue(line string, part int) (int, error) {
	firstDigit, lastDigit, err := extractDigits(line, part)
	if err != nil {
		return 0, fmt.Errorf("error extracting digits: %w", err)
	}
	return combineDigits(firstDigit, lastDigit), nil
}

// extractDigits returns the first and last digit of a line.
func extractDigits(line string, part int) (firstDigit, lastDigit int, err error) {
	switch part {
	case 1:
		return extractDigits1(line)
	case 2:
		return extractDigits2(line)
	default:
		return 0, 0, fmt.Errorf("invalid part")
	}
}

// extractDigits1 returns the first and last digit of a line written as digits.
func extractDigits1(line string) (firstDigit, lastDigit int, err error) {
	firstDigit, lastDigit = -1, -1
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			firstDigit = int(line[i] - '0')
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			lastDigit = int(line[i] - '0')
			break
		}
	}
	if firstDigit == -1 || lastDigit == -1 {
		return 0, 0, fmt.Errorf("no digits found")
	}
	return firstDigit, lastDigit, nil
}

var digits = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// extractDigits2 returns the first and last digit of a line written either as digit or with letters.
func extractDigits2(line string) (firstDigit, lastDigit int, err error) {
	firstDigit, lastDigit = -1, -1
	for i := 0; i < len(line) && firstDigit == -1; i++ {
		if line[i] >= '0' && line[i] <= '9' {
			firstDigit = int(line[i] - '0')
			break
		}
		for j := 0; j < len(digits); j++ {
			index_end := i + len(digits[j]) - 1
			if index_end < len(line) {
				subst := line[i : index_end+1]
				if subst == digits[j] {
					firstDigit = j
					break
				}
			}
		}
	}
	for i := len(line) - 1; i >= 0 && lastDigit == -1; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			lastDigit = int(line[i] - '0')
			break
		}
		for j := 0; j < len(digits); j++ {
			index_start := i - len(digits[j]) + 1
			if index_start >= 0 {
				subst := line[index_start : i+1]
				if subst == digits[j] {
					lastDigit = j
					break
				}
			}
		}
	}
	if firstDigit == -1 || lastDigit == -1 {
		return 0, 0, fmt.Errorf("no digits found")
	}
	return firstDigit, lastDigit, nil
}

// combineDigits returns an integer composed by the concatenation of two digits as bytes.
func combineDigits(firstDigit, lastDigit int) int {
	return firstDigit*10 + lastDigit
}
