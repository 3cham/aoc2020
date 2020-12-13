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
	bus := strings.Split(strings.ReplaceAll(lines[1], "x", "0"), ",")
	busId := utils.ConvertStringArrToInt64Arr(bus)
	return findEarliestPossibleTimestamp(busId)
}

func gcd(x, y int64) int64 {
	for x > 0 && y > 0 {
		if x > y {
			x %= y
		} else {
			y %= x
		}
	}
	if x > 0 {
		return x
	}
	return y
}

// find least common multiple
func lcm(x, y int64) int64 {
	return x * y / gcd(x, y)
}

func find(x, y, step, remain int64) int64 {
	current := int64(0)
	for {
		if (x+current)%y == remain {
			break
		}
		current += step
	}
	return x + current
}

func findEarliestPossibleTimestamp(busId []int64) int64 {
	current := busId[0]
	step := current
	for index, val := range busId {
		println(current, index, val)
		if index > 0 && val != 0 {
			current = find(current, val, step, val - int64(index))
			step = lcm(val, step)
		}
	}
	return current
}

func getFirstBus(content string) int {
	lines := strings.Split(content, "\n")
	timestamp, _ := strconv.Atoi(lines[0])
	bus := strings.ReplaceAll(lines[1], "x,", "")
	bus = strings.ReplaceAll(bus, "x", "")
	busId := utils.ConvertStringArrToIntArr(strings.Split(bus, ","))
	return findBusBasedOnTimestamp(timestamp, busId)
}

func findBusBasedOnTimestamp(timestamp int, busId []int) int {
	max := -1
	result := 0
	for _, bus := range busId {
		if timestamp%bus > max {
			max = timestamp % bus
			result = (bus - max) * bus
		} else if timestamp%bus == 0 {
			return 0
		}
	}
	return result
}
