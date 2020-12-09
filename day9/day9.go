package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

var markedPos []int

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day9/input.txt"
	content := utils.ReadInput(path)
	preamble := 25
	invalidNumber := findInvalidNumber(content, preamble)
	fmt.Println(invalidNumber)
	fmt.Println(findSumtoInvalid(content, invalidNumber))
}

func findSumtoInvalid(content string, invalidNumber int64) int64 {
	lines := strings.Split(content, "\n")
	numbers := utils.ConvertStringArrToIntArr(lines)
	start := 0
	end := 0
	for invalidNumber != 0 && end < len(lines) {
		if invalidNumber > 0 {
			invalidNumber -= numbers[end]
			end++
		} else if invalidNumber < 0 {
			invalidNumber += numbers[start]
			start++
		}
	}
	if invalidNumber == 0 {
		min, max := int64(1e9), -int64(1e9)
		for i := start; i < end; i++ {
			if numbers[i] > max {
				max = numbers[i]
			}
			if numbers[i] < min {
				min = numbers[i]
			}
		}
		return min + max
	}
	return -1
}

func findInvalidNumber(content string, preamble int) int64 {
	lines := strings.Split(content, "\n")
	numbers := utils.ConvertStringArrToIntArr(lines)

	for i := preamble; i < len(numbers); i++ {
		found := false
		for j := i - preamble; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if numbers[j] + numbers[k] == numbers[i] {
					found = true
				}
			}
			if found {
				break
			}
		}
		if !found {
			return numbers[i]
		}
	}
	return -1
}
