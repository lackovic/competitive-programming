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

	switch os.Args[1] {
	case "1":
		day01.Solve()
	default:
		fmt.Println("Unknown solution:", os.Args[1])
	}
}
