package main

import (
	"testing"
)

func TestReplaceEval(t *testing.T)  {
	testCases := []struct{
		expression string
		value, start, end int
		replacedExpression string
	}{
		{"1 + (2 * 3) + (4 * (5 + 6))", 6, 4, 10, "1 + 6 + (4 * (5 + 6))"},
		{"1 + (2 * 3) + (4 * (5 + 6))", 1, 14, 26, "1 + (2 * 3) + 1"},
	}
	for _, test := range testCases {
		t.Run("ReplaceEval should return correct result", func(t *testing.T) {
			got := replaceEval(test.expression, test.value, test.start, test.end)

			if got != test.replacedExpression {
				t.Fatalf("Wrong result, expected: %s, got: %s", test.replacedExpression, got)
			}
		})
	}
}


func TestFindDeepestParentheses(t *testing.T)  {
	testCases := []struct{
		expression string
		start, end int
	}{
		{"1 + (2 * 3) + (4 * 5 + 6)", 4, 10},
		{"1 + (2 * 3) + (4 * (5 + 6))", 19, 25},
		{"1 + 2 * 3 + 4 * 5 + 6", -1, -1},
	}
	for _, test := range testCases {
		t.Run("findDeepestParentheses should return correct result", func(t *testing.T) {
			start, end := findDeepestParentheses(test.expression)

			if start != test.start || end != test.end {
				t.Fatalf("Wrong result, expected: (%d, %d), got: (%d, %d)", test.start, test.end, start, end)
			}
		})
	}
}


func TestEvaluate(t *testing.T)  {
	testCases := []struct{
		expression string
		expected int
	}{
		{"1", 1},
		{"10", 10},
		{"10 + 1", 11},
		{"10 + 1 * 2", 22},
		{"10 + (1 * 2)", 12},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}
	for _, test := range testCases {
		t.Run("evaluate should return correct result", func(t *testing.T) {
			result := evaluate(test.expression)

			if result != test.expected {
				t.Fatalf("Wrong result, expected: %d, got: %d", test.expected, result)
			}
		})
	}
}