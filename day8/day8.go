package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strconv"
	"strings"
)

var markedPos []int

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day8/input.txt"
	content := utils.ReadInput(path)
	fmt.Println(findValueAfterChangedLine(content))
}

func getLastValueBeforeLoop(content string) int {
	lines := strings.Split(content, "\n")
	value := 0
	marked := make([]int, len(lines))

	currentLine := 0

	for ;; {
		if marked[currentLine] != 0 {
			return value
		}
		marked[currentLine] = 1
		value, currentLine = update(value, currentLine, lines[currentLine])
	}
}

func update(value int, position int, s string) (int, int) {
	ins, op, val := parseInstruction(s)
	switch ins {
	case "nop":
		return value, position + 1
	case "jmp":
		return value, cal(op, position, val)
	case "acc":
		return cal(op, value, val), position + 1
	default:
		return value, position
	}
}

func cal(op string, currentVal int, val int) int {
	switch op {
	case "-":
		return currentVal - val
	case "+":
		return currentVal + val
	default:
		return currentVal
	}
}

func parseInstruction(s string) (string, string, int) {
	words := strings.Split(s, " ")
	val, _ := strconv.Atoi(words[1][1:])

	return words[0], string(words[1][0]), val
}

func findValueAfterChangedLine(content string) int {
	lines := strings.Split(content, "\n")
	pos := 0
	markedPos = make([]int, len(lines))
	val, _ := findTerminatedValue(lines, pos, false, 0)
	return val
}

func findTerminatedValue(lines []string, pos int, changed bool, currentValue int) (int, bool) {
	if pos >= len(lines) {
		return currentValue, true
	}

	if markedPos[pos] > 0 {
		return -1, false
	}
	markedPos[pos] = 1
	oldLine := lines[pos]

	defer func(){
		markedPos[pos] = 0
		lines[pos] = oldLine
	}()
	ins, op, val := parseInstruction(lines[pos])
	switch ins {
	case "acc":
		return findTerminatedValue(lines, pos + 1, changed, cal(op, currentValue, val))
	case "jmp", "nop":
		_, newPos := update(currentValue, pos, lines[pos])
		newValue, found := findTerminatedValue(lines, newPos, changed, currentValue)
		if found {
			return newValue, found
		}

		if !changed {
			lines[pos] = switchIns(ins) + " " + op + string(val)
			_, newPos := update(currentValue, pos, lines[pos])
			newValue, found := findTerminatedValue(lines, newPos, true, currentValue)
			if found {
				return newValue, found
			}
		}
		return -1, false
	default:
		return -1, false
	}
}

func switchIns(s string) string {
	switch s{
	case "jmp":
		return "nop"
	default:
		return "jmp"
	}
}