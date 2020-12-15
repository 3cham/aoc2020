package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day15/input.txt"
	content := utils.ReadInput(path)
	fmt.Println("Result is: ", getNumber(content, 2020))

}

func getNumber(content string, turn int) int {
	numbers := strings.Split(content, ",")
	num := utils.ConvertStringArrToIntArr(numbers)
	lastPosition := make(map[int]int)

	current := 1
	for current < turn {
		lastKnownPosition, found := lastPosition[num[current - 1]]
		if current >= len(num) {
			if !found {
				num = append(num, 0)
			} else {
				num = append(num, current - 1 - lastKnownPosition)
			}
		}
		lastPosition[num[current - 1]] = current - 1
		//println(num[current])
		current++
	}
	return num[current - 1]
}
