package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strconv"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day18/input.txt"
	content := utils.ReadInput(path)
	fmt.Println("Result is: ", getSumExpressions(content))
}

func getSumExpressions(content string) int {
	lines := strings.Split(content, "\n")
	sum := 0
	for _, line := range lines {
		sum += evaluate(line)
	}
	return sum
}

func evaluate(line string) int {
	for {
		start, end := findDeepestParentheses(line)
		if start == -1 {
			break
		}
		value := evaluate(line[start + 1:end])
		line = replaceEval(line, value, start, end)
	}
	//TODO: evaluate without parentheses
	result := 0
	currentNum := 0
	line = strings.ReplaceAll(line, " ", "") + "+"
	lastOp := '+'
	for _, char := range line {
		switch char {
		case '+','*':
			switch lastOp {
			case '*':
				result *= currentNum
			default:
				result += currentNum
			}
			currentNum = 0
			lastOp = char
		default:
			digit, _ := strconv.Atoi(string(char))
			currentNum = currentNum * 10 + digit
		}
	}
	return result
}

func replaceEval(line string, value, start int, end int) string {
	return line[:start] + fmt.Sprint(value) + line[end+1:]
}

func findDeepestParentheses(line string) (int, int) {
	depth := 0
	maxDepth := 0
	start, end := -1, -1

	for index, char := range line {
		if char == '(' {
			depth += 1
			if depth > maxDepth {
				start = index
				maxDepth = depth
			}
		} else if char == ')' {
			depth -= 1
		}
	}

	for i := start + 1; i < len(line); i++ {
		if line[i] == ')' {
			end = i
			break
		}
	}

	return start, end
}
