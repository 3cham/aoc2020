package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strconv"
	"strings"
)

var rules map[int]string

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day19/input.txt"
	content := utils.ReadInput(path)
	fmt.Println("Result is: ", getValidStringCount(content))
}

func getValidStringCount(content string) int {
	splittedContent := strings.Split(content, "\n\n")
	rules = parseRules(splittedContent[0])
	count := 0
	lines := strings.Split(splittedContent[1], "\n")
	for _, line := range lines {
		if isMatched(line, rules[0]) == line {
			count++
		}
	}
	return count
}

func isMatched(line string, s string) string {
	if strings.Contains(s, "|") {
		subRules := strings.Split(s, " | ")
		maxLength := 0
		maxMatched := ""
		for _, rule := range subRules {
			matched := isMatched(line, rule)
			if len(matched) > maxLength {
				maxLength = len(matched)
				maxMatched = matched
			}
		}
		return maxMatched
	}
	if strings.Contains(s, "\"") {
		if strings.HasPrefix(line, strings.ReplaceAll(s, "\"", "")) {
			return strings.ReplaceAll(s, "\"", "")
		}
		return ""
	} else {
		ruleArr := utils.ConvertStringArrToIntArr(strings.Split(s, " "))
		currentMatch := line
		for _, rule := range ruleArr {
			part := isMatched(currentMatch, rules[rule])
			if part == "" {
				return part
			}
			currentMatch = currentMatch[len(part):]
		}
		return line[:len(line) - len(currentMatch)]
	}
	return line
}

func parseRules(s string) map[int]string {
	result := make(map[int]string)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		splittedLine := strings.Split(line, ": ")
		rulNo, _ := strconv.Atoi(splittedLine[0])
		result[rulNo] = splittedLine[1]
	}
	return result
}
