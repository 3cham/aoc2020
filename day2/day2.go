package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"path/filepath"
	"strings"
)

func main() {
	path, _ := filepath.Abs("/Users/tdang/go/src/github.com/3cham/aoc2020/day2/input.txt")
	content := utils.ReadInput(path)
	fmt.Println(countValidPassword(content))
}

func isValidPassword1(line string) bool{
	elem := strings.Split(line, ": ")
	word := elem[1]
	char := elem[0][len(elem[0]) - 1]
	rangeString := elem[0][:len(elem[0]) - 2]
	rangeCount := utils.ConvertStringArrToIntArr(strings.Split(rangeString, "-"))
	var count int64 = 0
	for i := 0; i < len(word); i++{
		if word[i] == char {
			count++
		}
	}
	return rangeCount[0] <= count && count <= rangeCount[1]
}

func isValidPassword2(line string) bool{
	elem := strings.Split(line, ": ")
	word := elem[1]
	char := elem[0][len(elem[0]) - 1]
	rangeString := elem[0][:len(elem[0]) - 2]
	rangeCount := utils.ConvertStringArrToIntArr(strings.Split(rangeString, "-"))
	return (word[rangeCount[0] - 1] == char || word[rangeCount[1] - 1] == char) && word[rangeCount[0] - 1] != word[rangeCount[1] - 1]
}

func countValidPassword(content string) int {
	println(content)
	validPassword := 0
	for _, line := range strings.Split(content, "\n") {
		if isValidPassword2(line) {
			validPassword++
		}
	}
	return validPassword
}
