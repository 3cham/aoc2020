package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

type Constraint struct {
	min, max int
}

func (c *Constraint) Satisfy(value int) bool {
	return c.min <= value && c.max >= value
}

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day16/input.txt"
	content := utils.ReadInput(path)
	fmt.Println("Result is: ", getSumInvalid(content))
}

func getSumInvalid(content string) int {
	lineBlock := strings.Split(content, "\n\n")
	constraints := parseConstraints(lineBlock[0])
	_ = parseMyTicket(lineBlock[1])
	otherTickets := parseOtherTickets(lineBlock[2])

	result := 0
	for _, ticket := range otherTickets {
		for _, val := range ticket {
			found := false
			for _, constraint := range constraints {
				if constraint.Satisfy(val) {
					found = true
					break
				}
			}
			if !found {
				result += val
			}
		}
	}
	return result
}

func parseOtherTickets(s string) [][]int {
	lines := strings.Split(s, "\n")[1:]
	result := [][]int{}
	for _, line := range lines {
		lineNum := strings.Split(line, ",")
		result = append(result, utils.ConvertStringArrToIntArr(lineNum))
	}
	return result
}

func parseMyTicket(s string) []int {
	line := strings.Split(s, "\n")[1]
	lineNum := strings.Split(line, ",")
	return utils.ConvertStringArrToIntArr(lineNum)
}

func parseConstraints(block string) []Constraint {
	result := []Constraint{}
	lines := strings.Split(block, "\n")
	for _, line := range lines {
		rangeValue := strings.Split(line, ": ")
		ranges := strings.Split(rangeValue[1], " or ")
		for _, rang := range ranges {
			rangNum := utils.ConvertStringArrToIntArr(strings.Split(rang, "-"))
			if rangNum[0] < rangNum[1] {
				result = append(result, Constraint{rangNum[0], rangNum[1]})
			} else {
				result = append(result, Constraint{rangNum[1], rangNum[0]})
			}
		}
	}
	return result
}


