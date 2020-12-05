package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day5/input.txt"
	content := utils.ReadInput(path)
	//println(findMaxSeatId(content))
	println(findFreeSeatId(content))
}

func findFreeSeatId(content string) int {
	lines := strings.Split(content, "\n")
	seated := [1023]int{}
	for _, line := range lines {
		seatId := calcSid(line)
		seated[seatId] = 1
	}
	for i := 1; i < 1023; i++ {
		if seated[i] == 0 {
			fmt.Printf("Free at %d\n", i)
		}
	}
	return -1
}

func findMaxSeatId(content string) int {
	lines := strings.Split(content, "\n")
	maxSid := 0
	for _, line := range lines {
		seatId := calcSid(line)
		println(seatId)
		if seatId > maxSid {
			maxSid = seatId
		}
	}
	return maxSid
}

func calcSid(line string) int {
	lr, hr, lc, hc := 0, 127, 0, 7
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case 'B':
			lr = (lr + hr + 1) / 2
		case 'F':
			hr = (lr + hr) / 2
		case 'R':
			lc = (lc + hc + 1) / 2
		case 'L':
			hc = (lc + hc) / 2
		}
	}
	return lr * 8 + lc
}
