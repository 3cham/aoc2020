package main

import (
	"testing"
)

func TestParseRules(t *testing.T)  {
	testCases := []struct{
		rule string
		ruleNo int
		expected string
	}{
		{"0: 4 1 5", 0, "4 1 5"},
		{"4: \"a\"", 4, "\"a\""},
		{"3: 4 5 | 5 4", 3, "4 5 | 5 4"},
	}
	for _, test := range testCases {
		t.Run("parseRules should return correct result", func(t *testing.T) {
			got := parseRules(test.rule)

			if got[test.ruleNo] != test.expected {
				t.Fatalf("Wrong result, expected: %d and %s", test.ruleNo, test.expected)
			}
		})
	}
}