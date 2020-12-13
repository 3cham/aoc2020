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
