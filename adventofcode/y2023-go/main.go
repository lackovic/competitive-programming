package main

import (
	"adventofcode2023/day01"
	"adventofcode2023/day02"
	"adventofcode2023/day03"
	"adventofcode2023/day04"
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
		result1, err1 = day01.Solve("day01/trebuchet.input", 1)
		result2, err2 = day01.Solve("day01/trebuchet.input", 2)
	case "2":
		result1, err1 = day02.Solve("day02/cube_conundrum.input")
		result2, err2 = day02.SolvePart2("day02/cube_conundrum.input")
	case "3":
		result1, err1 = day03.Solve1("day03/gear_ratios.input")
		result2, err2 = day03.Solve2("day03/gear_ratios.input")
	case "4":
		result1, err1 = day04.Solve1("day04/scratchcards.input")
		result2, err2 = day04.Solve2("day04/scratchcards.input")
	default:
		fmt.Println("Unknown solution:", os.Args[1])
	}
	if err1 != nil || err2 != nil {
		if err1 != nil {
			fmt.Println("Error1:", err1)
		}
		if err2 != nil {
			fmt.Println("Error2:", err2)
		}
		return
	}
	fmt.Println(result1, result2)
}
