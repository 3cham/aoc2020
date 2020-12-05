package main

import (
	"github.com/3cham/aoc2020/utils"
	"regexp"
	"strconv"
	"strings"
)

var fields = []string{
"byr",
"iyr",
"eyr",
"hgt",
"hcl",
"ecl",
"pid",
"cid",
}

func main() {
	path := "/Users/tdang/go/src/github.com/3cham/aoc2020/day4/input.txt"
	content := utils.ReadInput(path)
	println(countValidPassport(content))
}

func isValid(passport string) int {
	data := make(map[string]string)
	for _, line := range strings.Split(passport, "\n") {
		for _, entry := range strings.Split(line, " ") {
			kv := strings.Split(entry, ":")
			data[kv[0]] = kv[1]
		}
	}

	for _, key := range fields[:len(fields) - 1] {
		_, found := data[key]
		if !found {
			return 0
		}
	}
	return 1
}

func isValid2(passport string) int {
	data := make(map[string]string)
	for _, line := range strings.Split(passport, "\n") {
		for _, entry := range strings.Split(line, " ") {
			kv := strings.Split(entry, ":")
			data[kv[0]] = kv[1]
		}
	}

	for _, key := range fields[:len(fields) - 1] {
		_, found := data[key]
		if !found {
			return 0
		}
	}
	if !byrValid(data["byr"]) {
		return 0
	}
	if !iyrValid(data["iyr"]) {
		return 0
	}
	if !eyrValid(data["eyr"]) {
		return 0
	}
	if !hgtValid(data["hgt"]) {
		return 0
	}
	if !hclValid(data["hcl"]) {
		return 0
	}
	if !eclValid(data["ecl"]) {
		return 0
	}
	if !pidValid(data["pid"]) {
		return 0
	}
	return 1
}

func pidValid(s string) bool {
	pidCondition := regexp.MustCompile(`^[a-f0-9]{9}$`)
	return pidCondition.MatchString(s)
}

func eclValid(s string) bool {
	validVal := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, val := range validVal {
		if s == val {
			return true
		}
	}
	return false
}

func hclValid(s string) bool {
	nameCondition := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	return nameCondition.MatchString(s)
}

func hgtValid(s string) bool {
	v,err := strconv.Atoi(s[:len(s) - 2])
	if err != nil {
		return false
	}

	if strings.HasSuffix(s, "cm") {
		return v >= 150 && v <= 193
	} else if strings.HasSuffix(s,"in") {
		return v >= 59 && v <= 76
	}

	return false
}

func eyrValid(s string) bool {
	if len(s) != 4 {
		return false
	}
	yr, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return yr <= 2030 && yr >= 2020
}

func iyrValid(s string) bool {
	if len(s) != 4 {
		return false
	}
	yr, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return yr <= 2020 && yr >= 2010
}

func byrValid(s string) bool {
	if len(s) != 4 {
		return false
	}
	yr, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return yr <= 2002 && yr >= 1920
}

func countValidPassport(content string) int {
	passports := strings.Split(content, "\n\n")
	value := 0
	for _, passport := range passports {
		value += isValid2(passport)
	}
	return value
}
