package main

import (
	"fmt"
	"testing"
)

func TestGetNeighbor(t *testing.T) {
	c := Cube{0,0,0}
	t.Run("GetNeighbor should return 26 neighbors", func(t *testing.T) {
		neighbors := c.getNeighbors()
		if len(neighbors) != 26 {
			t.Fatalf("Wrong number of neighbors, got %v", len(neighbors))
		}
	})
	t.Run("test map of cube", func(t *testing.T) {
		active := make(map[Cube]bool)
		active[Cube{1,1,1}] = true
		for x := -1; x <= 1; x ++ {
			for y := -1; y <= 1; y ++ {
				for z := -1; z <= 1; z ++ {
					c := Cube{x,y,z}
					fmt.Printf("%v %v\n", c, active[c])
				}
			}
		}
		if active[Cube{0,0,0}] {
			t.Fatalf("Wrong value for 0,0,0")
		}
		if !active[Cube{1,1,1}] {
			t.Fatalf("Wrong value for 0,0,0")
		}
	})
}