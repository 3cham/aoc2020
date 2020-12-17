package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

type Cube struct {
	x, y, z int
}

func (c *Cube) getNeighbors() []Cube {
	result := []Cube{}
	for x := c.x - 1; x <= c.x + 1; x ++ {
		for y := c.y - 1; y <= c.y + 1; y ++ {
			for z := c.z - 1; z <= c.z + 1; z++ {
				if x != c.x || y != c.y || z != c.z {
					result = append(result, Cube{x, y, z})
				}
			}
		}
	}
	return result
}

var (
	active map[Cube]bool
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day17/input.txt"
	content := utils.ReadInput(path)
	activeCubes := parseCubeMap(content)
	fmt.Println(countActive(activeCubes, 6))
}

func countActive(activeCubes []Cube, count int) int {
	for count > 0 {
		activeCubes = updateState(activeCubes)
		count--
	}
	return len(activeCubes)
}

func updateState(activeCubes []Cube) []Cube {
	result := []Cube{}
	mapActiveCount := make(map[Cube]int)
	active := make(map[Cube]bool)
	for _, cube := range activeCubes {
		active[cube] = true
	}
	for _, cube := range activeCubes {
		neighbors := cube.getNeighbors()
		for _, neighbor := range neighbors {
			_, found := mapActiveCount[neighbor]
			if found {
				mapActiveCount[neighbor]++
			} else {
				mapActiveCount[neighbor] = 1
			}
		}
	}
	for _, cube := range activeCubes {
		if mapActiveCount[cube] == 2 || mapActiveCount[cube] == 3 {
			result = append(result, cube)
		}
	}
	for key, value := range mapActiveCount {
		if !active[key] && value == 3 {
			result = append(result, key)
		}
	}
	return result
}

func parseCubeMap(content string) []Cube {
	result := []Cube{}
	lines := strings.Split(content, "\n")
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				result = append(result, Cube{0, row, col})
			}
		}
	}
	return result
}