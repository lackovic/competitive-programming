package main

import (
	"adventofcode2023/day01"
	"adventofcode2023/day02"
	"adventofcode2023/day03"
	"adventofcode2023/day04"
	"adventofcode2023/day05"
	"adventofcode2023/day06"
	"adventofcode2023/day07"
	"adventofcode2023/day08"
	"adventofcode2023/day09"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide the solution day number")
		return
	}
	var result1, result2 int
	var result1_64, result2_64 int64
	var err1, err2 error
	switch os.Args[1] {
	case "1":
		result1, err1 = day01.Solve("day01/trebuchet.input", 1)
		result2, err2 = day01.Solve("day01/trebuchet.input", 2)
		fmt.Println(result1, result2)
	case "2":
		result1, err1 = day02.Solve("day02/cube_conundrum.input")
		result2, err2 = day02.SolvePart2("day02/cube_conundrum.input")
		fmt.Println(result1, result2)
	case "3":
		result1, err1 = day03.Solve1("day03/gear_ratios.input")
		result2, err2 = day03.Solve2("day03/gear_ratios.input")
		fmt.Println(result1, result2)
	case "4":
		result1, err1 = day04.Solve1("day04/scratchcards.input")
		result2, err2 = day04.Solve2("day04/scratchcards.input")
		fmt.Println(result1, result2)
	case "5":
		result1_64, err1 = day05.Solve1("day05/fertilizer.input")
		fmt.Println("Part 1 result = ", result1_64)
		result2_64, err2 = day05.Solve2("day05/fertilizer.input")
		fmt.Println("Part 2 result = ", result2_64)
	case "6":
		result1, err1 = day06.Solve1("day06/waitforit.input")
		fmt.Println("Part 1 result = ", result1)
		result2_64, err2 = day06.Solve2("day06/waitforit.input")
		fmt.Println("Part 2 result = ", result2_64)
	case "7":
		result1_64, err1 = day07.Solve1("day07/camel_cards.input")
		fmt.Println("Part 1 result = ", result1_64)
		result2_64, err2 = day07.Solve2("day07/camel_cards.input")
		fmt.Println("Part 2 result = ", result2_64)
	case "8":
		result1, err1 = day08.Solve1("day08/wasteland.input")
		result2_64, err2 = day08.Solve2("day08/wasteland.input")
		fmt.Println(result1, result2_64)
	case "9":
		result1, err1 = day09.Solve1("day09/mirage.input")
		fmt.Println("Part 1 result = ", result1)
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
}
