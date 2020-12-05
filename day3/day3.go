package main

import (
	"github.com/3cham/aoc2020/utils"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day3/input.txt"
	content := utils.ReadInput(path)
	checks := [][]int {{1,1},{1,3},{1,5},{1,7},{2,1}}
	result := 1
	for _, check := range checks {
		v := countTree(content, check[0], check[1])
		println(v)
		result *= v
	}
	println(result)
}

func check(lines []string, r,c int) (int, int, int) {
	if r >= len(lines) {
		return r,c, 0
	}
	r = r % len(lines)
	c = c % len(lines[0])
	if lines[r][c] == '#' {
		return r, c, 1
	} else {
		return r, c, 0
	}
}

func countTree(content string, r,c int) int {
	lines := strings.Split(content, "\n")
	start_r, start_c, count := check(lines, 0, 0)
	for start_r < len(lines) {
		v := 0
		start_r, start_c, v = check(lines, start_r + r, start_c + c)
		count += v
	}
	return count
}
