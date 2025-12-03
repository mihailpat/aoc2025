package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var timesZeroReached int
var currentPoint = 50

func Run() {
	InitSafeCracking("day1/input.txt")
}

func InitSafeCracking(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
	}
	content := string(data)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line != "" {
			direction := string(line[0])
			times, _ := strconv.Atoi(line[1:])
			currentPoint = dialTheVault(direction, times, currentPoint)

			if isZero(currentPoint) {
				timesZeroReached++
			}
		}

	}

	fmt.Println("Times zero reached: " + strconv.Itoa(timesZeroReached))
}

func dialTheVault(direction string, times int, currentPoint int) int {
	var tmpPosition int
	if direction == "L" {
		tmpPosition = currentPoint - times
	}
	if direction == "R" {
		tmpPosition = currentPoint + times
	}
	return tmpPosition % 100
}

func isZero(result int) bool {
	return result == 0
}
