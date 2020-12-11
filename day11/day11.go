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
	fmt.Println(countOccupied(change2(content)))
}

func countOccupied(s string) int {
	return strings.Count(s, "#")
}

func move(content string, countNeighborFunc func([]string, int, int) int, threshold int) string {
	lines := strings.Split(content, "\n")
	newContent := ""
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			x := countNeighborFunc(lines, i, j)
			if x == 0 && lines[i][j] == 'L' {
				newContent += "#"
				continue
			}
			if x >= threshold && lines[i][j] == '#' {
				newContent += "L"
				continue
			}
			newContent += string(lines[i][j])
		}
		if i < len(lines) - 1 {
			newContent += "\n"
		}
	}
	return newContent
}

func countNeighbor(lines []string, i,j int) int {
	neighbors := []struct {
		x int
		y int
	}{
		{-1,-1},
		{-1,0},
		{-1,1},
		{0,-1},
		{0,1},
		{1,-1},
		{1,0},
		{1,1},
	}
	count := 0
	for _, neighbor := range neighbors {
		ii := i + neighbor.x
		jj := j + neighbor.y

		if ii >= 0 && ii < len(lines) && jj >= 0 && jj < len(lines[i]) {
			if lines[ii][jj] == '#' {
				count++
			}
		}
	}
	return count
}

func countNeighborInDirection(lines []string, i,j int) int {
	neighbors := []struct {
		x int
		y int
	}{
		{-1,-1},
		{-1,0},
		{-1,1},
		{0,-1},
		{0,1},
		{1,-1},
		{1,0},
		{1,1},
	}
	count := 0

	for _, neighbor := range neighbors {
		a,b := i,j

		for {
			ii,jj := a + neighbor.x, b + neighbor.y
			if ii >= 0 && ii < len(lines) && jj >= 0 && jj < len(lines[i]) {
				if lines[ii][jj] == '#' {
					count++
					break
				}
				if lines[ii][jj] == 'L' {
					break
				}
			} else {
				break
			}
			a, b = ii,jj
		}
	}
	return count
}

func change(content string) string {
	for {
		newContent := move(content, countNeighbor, 4)
		if newContent == content {
			return newContent
		} else {
			content = newContent
		}
	}
	return content
}


func change2(content string) string {
	for {
		newContent := move(content, countNeighborInDirection, 5)
		if newContent == content {
			return newContent
		} else {
			content = newContent
		}
	}
	return content
}