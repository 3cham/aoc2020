package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strconv"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day14/input.txt"
	content := utils.ReadInput(path)
	fmt.Println(getSum(content))
	fmt.Println(getFloatingSum(content))
}

func getFloatingSum(content string) int64 {
	lines := strings.Split(content, "\n")
	memory := make(map[int64]int64)
	currentMask := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			currentMask = strings.Split(line, " = ")[1]
		} else {
			address, value := parseAddressAndValue(line)
			newAddresses := updateAddress(address, currentMask)
			for _, addr := range newAddresses {
				memory[addr] = value
			}
		}
	}
	result := int64(0)
	for _, value := range memory {
		result += value
	}
	return result
}

func updateAddress(address int64, mask string) []int64 {
	newAddresses := []int64 {address}
	currentP := int64(1)
	for i := len(mask) - 1; i >= 0; i-- {
		switch mask[i] {
		case '1':
			for index, value := range newAddresses {
				newAddresses[index] = value | currentP
			}
		case 'X':
			generatedAddr := []int64{}
			for _, value := range newAddresses {
				newValue0 := value & ^(currentP)
				newValue1 := value | currentP
				if value != newValue0 {
					generatedAddr = append(generatedAddr, newValue0)
				} else {
					generatedAddr = append(generatedAddr, newValue1)
				}
			}
			for _, value := range generatedAddr {
				newAddresses = append(newAddresses, value)
			}
		}
		currentP *= 2
	}
	return newAddresses
}

func parseAddressAndValue(line string) (int64, int64) {
	assignment := strings.Split(line, " = ")
	address, _ := strconv.Atoi(assignment[0][4:len(assignment[0]) - 1])
	value, _ := strconv.Atoi(assignment[1])
	return int64(address), int64(value)
}

func getSum(content string) int64 {
	lines := strings.Split(content, "\n")
	memory := make(map[int64]int64)
	currentMask := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			currentMask = strings.Split(line, " = ")[1]
		} else {
			address, value := parseAddressAndValue(line)
			value = updateValue(value, currentMask)
			memory[address] = value
		}
	}
	result := int64(0)
	for _, value := range memory {
		result += value
	}
	return result
}

func updateValue(value int64, mask string) int64 {
	currentP := int64(1)
	for i := len(mask) - 1; i >= 0; i-- {
		switch mask[i] {
		case '0':
			value = value & (^currentP)
		case '1':
			value = value | currentP
		}
		currentP *= 2
	}
	return value
}
