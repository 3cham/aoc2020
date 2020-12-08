package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strconv"
	"strings"
)

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day7/input.txt"
	content := utils.ReadInput(path)
	fmt.Println(countBagPart1(content))
	fmt.Println(countBagPart2(content) - 1)
}

func countBagPart1(content string) int {
	containMap := make(map[string][]string)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		outerBag, innerBags := parse(line)
		for _, bag := range innerBags {
			_, found := containMap[bag]
			if !found {
				containMap[bag] = []string{ outerBag }
			} else {
				containMap[bag] = append(containMap[bag], outerBag)
			}
		}
	}
	return countParrent(containMap, "shiny gold")
}

func countParrent(containMap map[string][]string, s string) int {
	marked := make(map[string]int)
	queue := []string{s}
	marked[s] = 1

	parent := 0
	for len(queue) != 0 {
		visited := queue[0]
		if visited != s {
			parent++
		}
		queue = queue[1:]
		for _, next := range containMap[visited] {
			_, found := marked[next]
			if ! found {
				queue = append(queue, next)
				marked[next] = 1
			}
		}
	}
	return parent
}

func parse(line string) (string, []string) {
	relation := strings.Split(line, " contain ")
	outerBag := relation[0][:len(relation[0]) - len(" bags")]
	if !strings.HasPrefix(relation[1], "no other") {
		innerBags := []string{}
		bagStrings := strings.Split(relation[1], ", ")

		for _, bagString := range bagStrings {
			words := strings.Split(bagString, " ")
			innerBags = append(innerBags, words[1] + " " + words[2])
		}
		return outerBag, innerBags
	}

	return outerBag, []string{}
}


func parse_with_count(line string) (string, map[string]int) {
	relation := strings.Split(line, " contain ")
	outerBag := relation[0][:len(relation[0]) - len(" bags")]
	result := make(map[string]int)
	if !strings.HasPrefix(relation[1], "no other") {
		bagStrings := strings.Split(relation[1], ", ")
		for _, bagString := range bagStrings {
			words := strings.Split(bagString, " ")
			innerBag := words[1] + " " + words[2]
			number, _ := strconv.Atoi(words[0])
			result[innerBag] = number
		}
	}

	return outerBag, result
}

func countBagPart2(content string) int {
	containMap := make(map[string]map[string]int)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		outerBag, innerBags := parse_with_count(line)
		containMap[outerBag] = innerBags
	}
	return countChildren(containMap, "shiny gold")
}

func countChildren(containMap map[string]map[string]int, s string) int {
	children := 1
	for next, val := range containMap[s] {
		children += val * countChildren(containMap, next)
	}
	return children
}
