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

*/

package day01

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Solve(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()
	calibrationDocument, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	lines := strings.Split(string(calibrationDocument), "\n")
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		calibrationVal, err := calibrationValue(line)
		if err != nil {
			fmt.Println("Error calculating calibration value:", err)
			return 0
		}
		total += calibrationVal
	}
	return total
}

// calibrationValue returns the calibration value of a line.
func calibrationValue(line string) (int, error) {
	firstDigit, lastDigit, err := extractDigits(line)
	if err != nil {
		return 0, fmt.Errorf("error extracting digits: %w", err)
	}
	return combineDigits(firstDigit, lastDigit), nil
}

// extractDigits returns the first and last digit of a line.
func extractDigits(line string) (firstDigit, lastDigit byte, err error) {
	firstIndex, lastIndex := -1, -1
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			firstIndex = i
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			lastIndex = i
			break
		}
	}
	if firstIndex == -1 || lastIndex == -1 {
		return 0, 0, fmt.Errorf("no digits found")
	}
	return line[firstIndex], line[lastIndex], nil
}

// combineDigits returns an integer composed by the concatenation of two digits as bytes.
func combineDigits(firstDigit, lastDigit byte) int {
	return int(firstDigit-'0')*10 + int(lastDigit-'0')
}
