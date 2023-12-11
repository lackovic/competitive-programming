package main

import (
	"adventofcode2023/day01"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide the solution day number")
		return
	}
	var result, result2 int
	switch os.Args[1] {
	case "1":
		result = day01.Solve("day01/trebuchet.input", 1)
		result2 = day01.Solve("day01/trebuchet.input", 2)
	default:
		fmt.Println("Unknown solution:", os.Args[1])
	}
	fmt.Println(result, result2)
}
