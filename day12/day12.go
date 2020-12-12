package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day12/input.txt"
	content := utils.ReadInput(path)
	fmt.Println(manhattanDistance(content))

}

func manhattanDistance(content string) float64 {
	moves := strings.Split(content, "\n")

	x,y, currentAngle := 0.0, 0.0, 0.0

	for _, move := range moves {
		x,y,currentAngle = update(x,y,currentAngle,move)
	}
	return math.Abs(x) + math.Abs(y)
}

func update(x float64, y float64, angle float64, move string) (float64, float64, float64) {
	ins := move[:1]
	value, _ := strconv.Atoi(move[1:])
	distance := float64(value)
	switch ins {
	case "F":
		x, y = updatePosition(x, y, angle, distance)
		return x, y, angle
	case "N":
		return x, y + distance, angle
	case "S":
		return x, y - distance, angle
	case "E":
		return x + distance, y, angle
	case "W":
		return x - distance, y, angle
	case "L":
	default:

	}
}

func updatePosition(x float64, y float64, angle float64, distance float64) (float64, float64) {

}
