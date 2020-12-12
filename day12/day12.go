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
	fmt.Println(manhattanDistance2(content))
}

func manhattanDistance2(content string) float64 {
	moves := strings.Split(content, "\n")

	xS,yS, xW, yW := 0.0, 0.0, 10.0, 1.0

	for _, move := range moves {
		xS,yS,xW, yW = updateSW(xS,yS,xW,yW,move)
	}
	return math.Abs(xS) + math.Abs(yS)
}

func updateSW(xS, yS, xW, yW float64, move string) (float64, float64, float64, float64) {
	ins := move[:1]
	value, _ := strconv.Atoi(move[1:])
	distance := float64(value)

	switch ins {
	case "F":
		return xS + xW * distance, yS + yW * distance, xW, yW
	case "N":
		return xS, yS, xW, yW + distance
	case "S":
		return xS, yS, xW, yW - distance
	case "E":
		return xS, yS, xW + distance, yW
	case "W":
		return xS, yS, xW - distance, yW
	case "L":
		distance = distance / 180 * math.Pi
		return xS, yS, xW * math.Cos(distance) - yW * math.Sin(distance), xW * math.Sin(distance) + yW * math.Cos(distance)
	default:
		distance = (360 - distance) / 180 * math.Pi
		return xS, yS, xW * math.Cos(distance) - yW * math.Sin(distance), xW * math.Sin(distance) + yW * math.Cos(distance)
	}
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
		return x + distance, y, angle
	case "S":
		return x - distance, y, angle
	case "E":
		return x, y + distance, angle
	case "W":
		return x, y - distance, angle
	case "L":
		angle = float64((int(angle) + int(distance)) % 360)
		return x, y, angle
	default:
		angle = float64((int(angle) - int(distance) + 360) % 360)
		return x, y, angle
	}
}

func updatePosition(x float64, y float64, angle float64, distance float64) (float64, float64) {
	return x + math.Sin(angle / 180 * math.Pi) * distance, y + math.Cos(angle / 180 * math.Pi) * distance
}
