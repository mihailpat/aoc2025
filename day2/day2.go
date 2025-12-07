package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	ranges := fileInput("day2/input.txt")
	checkValidIDs1(ranges)
}

func fileInput(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
	}
	content := string(data)
	ranges := strings.Split(content, ",")
	return ranges
}

func checkValidIDs1(ranges []string) {
	sum := 0
	for _, r := range ranges {
		subRange := strings.Split(r, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(subRange[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(subRange[1]))
		for i := start; i <= end; i++ {
			nString := strconv.Itoa(i)
			if len(nString)%2 != 0 {
				continue
			}

			mid := len(nString) / 2
			left := nString[:mid]
			right := nString[mid:]

			if left == right {
				nInt, _ := strconv.Atoi(nString)
				sum += nInt
			}
		}
	}
	fmt.Println(sum)
}
