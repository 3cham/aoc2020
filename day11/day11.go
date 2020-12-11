package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day11/input.txt"
	content := utils.ReadInput(path)
	fmt.Println(countOccupied(change(content)))
}

func countOccupied(s string) int {
	return strings.Count(s, "#")
}

func change(content string) string {
	return ""
}