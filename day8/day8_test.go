package main

import (
	"reflect"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	type Parsed struct {
		ins string
		op string
		val int
	}

	testCases := []struct {
		name string
		s string
		expected Parsed
	}{
		{"parseInstruction should parse noop", "nop +0", Parsed{"nop", "+", 0}},
		{"parseInstruction should parse jmp", "jmp -1", Parsed{"jmp", "-", 1}},
		{"parseInstruction should parse acc", "acc +10", Parsed{"acc", "+", 10}},

	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ins, op, val := parseInstruction(test.s)
			got := Parsed{ins, op, val}
			if !reflect.DeepEqual(got, test.expected) {
				t.Fatalf("Wrong values are parsed, got %v, want %v", got, test.expected)
			}
		})
	}
}
