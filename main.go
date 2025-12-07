package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mihailpat/aoc2025/day1"
	"github.com/mihailpat/aoc2025/day2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a day number to run (e.g., 'go run main.go 1')")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid day number: %s\n", os.Args[1])
		return
	}

	switch day {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}
}
