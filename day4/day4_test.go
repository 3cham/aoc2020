package main

import "testing"

func TestByrValid(t *testing.T) {
	testCases := []struct{
		yr string
		valid bool
	}{
		{"1920", true},
		{"2020", false},
		{"01940", false},
	}
	for _, tcase := range testCases {
		if byrValid(tcase.yr) != tcase.valid {
			t.Fatalf("Wrong result with %s, expected %v", tcase.yr, tcase.valid)
		}
	}
}

func TestIyrValid(t *testing.T) {
	testCases := []struct{
		yr string
		valid bool
	}{
		{"1920", false},
		{"2020", true},
		{"01940", false},
	}
	for _, tcase := range testCases {
		if iyrValid(tcase.yr) != tcase.valid {
			t.Fatalf("Wrong result with %s, expected %v", tcase.yr, tcase.valid)
		}
	}
}

func TestEyrValid(t *testing.T) {
	testCases := []struct{
		yr string
		valid bool
	}{
		{"1920", false},
		{"2020", true},
		{"01940", false},
		{"2030", true},
		{"2031", false},
	}
	for _, tcase := range testCases {
		if eyrValid(tcase.yr) != tcase.valid {
			t.Fatalf("Wrong result with %s, expected %v", tcase.yr, tcase.valid)
		}
	}
}

func TestHgtValid(t *testing.T) {
	testCases := []struct{
		yr string
		valid bool
	}{
		{"1920", false},
		{"2020", false},
		{"01940in", false},
		{"2030cm", false},
		{"x2030cm", false},
		{"02030cm", false},
		{"2031", false},
		{"174cm", true},
		{"78in", false},
	}
	for _, tcase := range testCases {
		if hgtValid(tcase.yr) != tcase.valid {
			t.Fatalf("Wrong result with %s, expected %v", tcase.yr, tcase.valid)
		}
	}
}


func TestHclValid(t *testing.T) {
	testCases := []struct{
		yr string
		valid bool
	}{
		{"#1920", false},
		{"2020", false},
		{"#019401", true},
		{"#2030cm", false},
		{"#2030af", true},
	}
	for _, tcase := range testCases {
		if hclValid(tcase.yr) != tcase.valid {
			t.Fatalf("Wrong result with %s, expected %v", tcase.yr, tcase.valid)
		}
	}
}


func TestEclValid(t *testing.T) {
	testCases := []struct{
		yr string
		valid bool
	}{
		{"#1920", false},
		{"2020", false},
		{"amb", true},
		{"blu", true},
		{"brn", true},
		{"gry", true},
		{"grn", true},
		{"hzl", true},
		{"oth", true},
	}
	for _, tcase := range testCases {
		if eclValid(tcase.yr) != tcase.valid {
			t.Fatalf("Wrong result with %s, expected %v", tcase.yr, tcase.valid)
		}
	}
}


func TestPidValid(t *testing.T) {
	testCases := []struct{
		yr string
		valid bool
	}{
		{"#1920", false},
		{"2020", false},
		{"487827amb", false},
		{"012345678", true},
		{"698275677", true},
		{"g242422ry", false},
	}
	for _, tcase := range testCases {
		if pidValid(tcase.yr) != tcase.valid {
			t.Fatalf("Wrong result with %s, expected %v", tcase.yr, tcase.valid)
		}
	}
}