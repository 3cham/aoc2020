package main

import "testing"

func TestLcm(t *testing.T) {
	testCases := []struct {
		x int64
		y int64
		lcd int64
	}{
		{2,4,4},
		{2,3,6},
		{4,6,12},
	}
	for _, test := range testCases {
		t.Run("LCM should return correct result", func(t *testing.T) {
			got := lcm(test.x, test.y)
			if got != test.lcd {
				t.Fatalf("Wrong result expect %d, got %d", test.lcd, got)
			}
		})
	}
}