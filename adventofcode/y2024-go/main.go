package main

import (
	"adventofcode2024/day01"
	"adventofcode2024/day02"
	"adventofcode2024/day03"
	"adventofcode2024/day04"
	"adventofcode2024/day05"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide the solution day number")
		return
	}
	var result1, result2 int
	var err1, err2 error
	switch os.Args[1] {
	case "1":
		result1, err1 = day01.Part1("day01/input")
		fmt.Println("Solution 1: ", result1)
		result2, err2 = day01.Part2("day01/input")
		fmt.Println("Solution 2: ", result2)
	case "2":
		result1, err1 = day02.Part1("day02/input")
		fmt.Println("Solution 1: ", result1)
		result2, err2 = day02.Part2("day02/input")
		fmt.Println("Solution 2: ", result2)
	case "3":
		result1, err1 = day03.Part1("day03/input")
		fmt.Println("Solution 1: ", result1)
		result2, err2 = day03.Part2("day03/input")
		fmt.Println("Solution 2: ", result2)
	case "4":
		result1, err1 = day04.Part1("day04/input")
		fmt.Println("Solution 1: ", result1)
		result2, err2 = day04.Part2("day04/input")
		fmt.Println("Solution 2: ", result2)
	case "5":
		result1, err1 = day05.Part1("day05/input")
		fmt.Println("Solution 1: ", result1)
		result2, err2 = day05.Part2("day05/input")
		fmt.Println("Solution 2: ", result2)
	default:
		fmt.Println("Unknown solution:", os.Args[1])
	}
	if err1 != nil {
		fmt.Println("Error1:", err1)
	}
	if err2 != nil {
		fmt.Println("Error2:", err2)
	}
	return
}
