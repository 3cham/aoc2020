package main

import (
	"fmt"
	"github.com/3cham/aoc2020/utils"
	"path/filepath"
	"strings"
)

func findProduct(numbers []int64) int64 {
	fmt.Println(numbers)
	for _, num := range numbers {
		for _, num2 := range numbers {
			if num + num2 == 2020 {
				return num * num2
			}
		}
	}
	return -1
}

func find3Product(numbers []int64) int64 {
	for _, num := range numbers {
		for _, num2 := range numbers {
			for _, num3 := range numbers {
				if num+num2 +num3== 2020 {
					return num * num2 * num3
				}
			}
		}
	}
	return -1
}

func main()  {
	path, _ := filepath.Abs("./input.txt")
	content := utils.ReadInput(path)
	numbers := utils.ConvertStringArrToInt64Arr(strings.Split(content, "\n"))
	println(findProduct(numbers))
	println(find3Product(numbers))
}
