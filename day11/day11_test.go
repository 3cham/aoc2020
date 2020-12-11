package main

import (
	"fmt"
	"testing"
)

func TestCountOccupied(t *testing.T) {
	testCases := []struct {
		value string
		expectedCount int
	}{
		{"#.##.##.##", 7},
		{"#######.##", 9},
		{"#.#.#..#..", 4},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("should return correct count for %s", testCase.value), func(t *testing.T) {
			got := countOccupied(testCase.value)
			if got != testCase.expectedCount {
				t.Fatalf("Expected: %d, got %d", testCase.expectedCount, got)
			}
		})
	}
}

func TestCountNeighbors(t *testing.T) {
	lines := []string {
		"#.##.##.##",
		"#######.##",
		"#.#.#..#..",
	}

	testCases := []struct {
		x, y, count int
	}{
		{0,0, 2},
		{1,1,6},
		{2,2,3},
	}

	for _, testCase := range testCases {
		t.Run("countNeighbors() should return correct number of occupied seats", func(t *testing.T) {
			got := countNeighbor(lines, testCase.x, testCase.y)
			if got != testCase.count {
				t.Fatalf("Incorrect number of occupied seats, got %d, expected %d", got, testCase.count)
			}
		})
	}
}


func TestCountNeighborInDirection(t *testing.T) {
	lines := []string {
		"#.##..#.L.",
		"######...L",
		"#.#.#..#..",
	}

	testCases := []struct {
		x, y, count int
	}{
		{1,9, 1},
		{2,9,1},
		{2,7,1},
	}

	for _, testCase := range testCases {
		t.Run("countNeighborIndirection() should return correct number of occupied seats", func(t *testing.T) {
			got := countNeighborInDirection(lines, testCase.x, testCase.y)
			if got != testCase.count {
				t.Fatalf("Incorrect number of occupied seats, got %d, expected %d", got, testCase.count)
			}
		})
	}
}
func TestMove(t *testing.T) {
	content := 	"#.##.##.##\n"+
				"#######.##\n"+
				"#.#.#..#.."
	got := move(content, countNeighbor, 4)
	expected := "#.LL.L#.##\n"+
				"#LLLLLL.L#\n"+
				"#.#.#..#.."

	if got != expected {
		t.Fatalf("Wrong move, got \n%s, expected \n%s", got, expected)
	}
}