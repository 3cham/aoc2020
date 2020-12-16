package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

type Constraint struct {
	name string
	min, max int
}

var (
	suitable map[string]map[int]bool
	used map[string]bool
)


func (c *Constraint) Satisfy(value int) bool {
	return c.min <= value && c.max >= value
}

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day16/input.txt"
	content := utils.ReadInput(path)
	fmt.Println("Result is: ", getSumInvalid(content))
	fmt.Println(getProductMyTicket(content))
}

func getProductMyTicket(content string) int {
	lineBlock := strings.Split(content, "\n\n")
	constraints := parseConstraints(lineBlock[0])
	constraintsMap := makeConstraintMap(constraints)
	myTicket := parseMyTicket(lineBlock[1])
	otherTickets := parseOtherTickets(lineBlock[2])
	validTickets := [][]int{}
	for _, ticket := range otherTickets {
		if isValidTicket(ticket, constraints) {
			validTickets = append(validTickets, ticket)
		}
	}
	validTickets = append(validTickets, myTicket)

	initializeMap(constraintsMap, validTickets)
	constraintsNames := searchConstraintsNames([]string{}, constraintsMap, validTickets)
	result := 1
	for index, name := range constraintsNames {
		if strings.HasPrefix(name, "departure") {
			result *= myTicket[index]
		}
	}
	return result
}

func initializeMap(constraintsMap map[string][]Constraint, tickets [][]int) {
	// initialize used map
	used = make(map[string]bool)
	for key, _ := range constraintsMap {
		used[key] = false
	}

	// pre-check if a constraint is suitable for a position or not
	suitable = make(map[string]map[int]bool)
	for key, value := range constraintsMap {
		suitable[key] = make(map[int]bool)
		for index, _ := range tickets[0] {
			suitable[key][index] = satisfyConstraintName(value, index, tickets)
		}
	}
}

func alreadyUsed(name string, currentNames []string) bool {
	for _, val := range currentNames {
		if val == name {
			return true
		}
	}
	return false
}

func searchConstraintsNames(currentNames []string, constraintsMap map[string][]Constraint, tickets [][]int) []string {

	if len(currentNames) == len(constraintsMap) {
		return currentNames
	}

	for key, _ := range constraintsMap {
		if !used[key] {
			if suitable[key][len(currentNames)] {
				used[key] = true
				foundNames := searchConstraintsNames(append(currentNames, key), constraintsMap, tickets)
				if foundNames != nil {
					return foundNames
				}
				used[key] = false
			}
		}
	}
	return nil
}

func satisfyConstraintName(cons []Constraint, i int, tickets [][]int) bool {
	for _, ticket := range tickets {
		if !cons[0].Satisfy(ticket[i]) && !cons[1].Satisfy(ticket[i]) {
			return false
		}
	}
	return true
}

func makeConstraintMap(constraints []Constraint) map[string][]Constraint {
	result := make(map[string][]Constraint)
	for _, cons := range constraints {
		val, found := result[cons.name]
		if !found {
			result[cons.name] = []Constraint{cons}
		} else {
			result[cons.name] = append(val, cons)
		}
	}
	return result
}

func isValidTicket(ticket []int, constraints []Constraint) bool {
	for _, val := range ticket {
		found := false
		for _, constraint := range constraints {
			if constraint.Satisfy(val) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
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
		name := rangeValue[0]
		ranges := strings.Split(rangeValue[1], " or ")
		for _, rang := range ranges {
			rangNum := utils.ConvertStringArrToIntArr(strings.Split(rang, "-"))
			if rangNum[0] < rangNum[1] {
				result = append(result, Constraint{name, rangNum[0], rangNum[1]})
			} else {
				result = append(result, Constraint{name, rangNum[1], rangNum[0]})
			}
		}
	}
	return result
}


