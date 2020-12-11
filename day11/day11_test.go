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