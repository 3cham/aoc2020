package main

import (
	"github.com/3cham/aoc2020/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day10/input.txt"
	content := utils.ReadInput(path)
	fmt.Println(findDist(content))
	fmt.Println(findArrangementCount(content))
}

func findArrangementCount(content string) int {
	lines := strings.Split(content, "\n")
	numbers := utils.ConvertStringArrToIntArr(lines)
	numbers = append(numbers, 0)

	sort.Ints(numbers)
	count := make([]int, len(numbers))
	count[0] = 1
	for i := 1; i < len(numbers); i++ {
		for j := i - 1; j >= 0; j-- {
			if numbers[i] - numbers[j] <= 3 {
				count[i] += count[j]
			} else {
				break
			}
		}
	}
	return count[len(numbers) - 1]
}

func findDist(content string) int {
	lines := strings.Split(content, "\n")
	numbers := utils.ConvertStringArrToIntArr(lines)
	numbers = append(numbers, 0)

	sort.Ints(numbers)
	count1 := 0
	count3 := 0
	for i := 1; i < len(numbers); i++ {
		switch numbers[i] - numbers[i - 1] {
		case 1:
			count1++
		case 3:
			count3++
		}
	}
	return count1 * (count3 + 1)
}

