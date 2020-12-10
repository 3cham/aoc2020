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