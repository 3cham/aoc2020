package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"strings"
)

func main()  {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day6/input.txt"
	content := utils.ReadInput(path)
	fmt.Println(countAnswerPart2(content))
}

func countAnswer(content string) int {
	groups := strings.Split(content, "\n\n")
	count := 0
	for _, group := range groups {
		answerCount := make(map[uint8]int)
		answers := strings.Split(group, "\n")
		for _, person := range answers {
			for i := 0; i < len(person); i++ {
				_, found := answerCount[person[i]]
				if !found {
					count++
					answerCount[person[i]] = 1
				}
			}
		}
	}
	return count
}


func countAnswerPart2(content string) int {
	groups := strings.Split(content, "\n\n")
	count := 0
	for _, group := range groups {
		answerCount := make(map[uint8]int)
		answers := strings.Split(group, "\n")
		for _, person := range answers {
			for i := 0; i < len(person); i++ {
				_, found := answerCount[person[i]]
				if !found {
					answerCount[person[i]] = 1
				} else {
					answerCount[person[i]] ++
				}
			}
		}
		for _,v := range answerCount {
			if v == len(answers) {
				count++
			}
		}
	}
	return count
}
