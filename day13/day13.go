package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strconv"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day13/input.txt"
	content := utils.ReadInput(path)
	fmt.Println(getFirstBus(content))
	fmt.Println(getFirstTimestamp(content))
}

func getFirstTimestamp(content string) int64 {
	lines := strings.Split(content, "\n")
	bus := strings.Split(strings.ReplaceAll(lines[1], "x","0"), ",")
	busId := utils.ConvertStringArrToInt64Arr(bus)
	return findEarliestPossibleTimestamp(busId)
}

// find least common multiple
func lcm(x, y int64) int64 {
	result := x * y
	for x > 0 && y > 0 {
		if x > y {
			x = x % y
		} else {
			y = y % x
		}
	}
	if x > 0 {
		return result / x
	} else {
		return result / y
	}
}

func findEarliestPossibleTimestamp(busId []int64) int64 {
	//for _, val := range busId {
	//
	//}
	return 0
}

func getFirstBus(content string) int {
	lines := strings.Split(content, "\n")
	timestamp, _ := strconv.Atoi(lines[0])
	bus := strings.ReplaceAll(lines[1], "x,","")
	bus = strings.ReplaceAll(bus, "x", "")
	busId := utils.ConvertStringArrToIntArr(strings.Split(bus, ","))
	return findBusBasedOnTimestamp(timestamp, busId)
}

func findBusBasedOnTimestamp(timestamp int, busId []int) int {
	max := -1
	result := 0
	for _, bus := range busId {
		if timestamp % bus > max {
			max = timestamp % bus
			result = (bus - max) * bus
		} else if timestamp % bus == 0 {
			return 0
		}
	}
	return result
}
