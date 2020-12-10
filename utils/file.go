package utils

import (
	"io/ioutil"
	"strconv"
)

func ReadInput(filename string) string {
	content, _ := ioutil.ReadFile(filename)
	return string(content)
}

func ConvertStringArrToInt64Arr(numbers []string) []int64 {
	convertedNumber := []int64{}

	for _, num := range numbers {
		numb, _ := strconv.Atoi(num)
		convertedNumber = append(convertedNumber, int64(numb))
	}
	return convertedNumber
}

func ConvertStringArrToIntArr(numbers []string) []int {
	convertedNumber := []int{}

	for _, num := range numbers {
		numb, _ := strconv.Atoi(num)
		convertedNumber = append(convertedNumber, numb)
	}
	return convertedNumber
}
