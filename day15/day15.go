package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day15/input.txt"
	content := utils.ReadInput(path)
	fmt.Println("Result is: ", getNumber(content, 30000000))
}

func getNumber(content string, turn int) int {
	numbers := strings.Split(content, ",")
	num := utils.ConvertStringArrToIntArr(numbers)
	lastPosition := make(map[int]int)

	current := 1
	lastNumber := num[0]
	for current < turn {
		lastKnownPosition, found := lastPosition[lastNumber]
		previousNumber := lastNumber
		if current >= len(num) {
			if !found {
				lastNumber = 0
			} else {
				lastNumber = current - 1 - lastKnownPosition
			}
		} else {
			lastNumber = num[current]
		}
		lastPosition[previousNumber] = current - 1
		//println(num[current])
		current++
	}
	return lastNumber
}
