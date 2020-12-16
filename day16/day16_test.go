package main

import (
	"reflect"
	"testing"
)

func checkIn(t *testing.T, c Constraint, constraints []Constraint) {
	t.Helper()
	found := false
	for _, cons := range constraints {
		if cons.name == c.name && cons.min == c.min && cons.max == c.max {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Constraints array: %v does not contain %v", constraints, c)
	}
}
func TestParseConstraints(t *testing.T)  {
	t.Run("parseConstraint should return correct result", func(t *testing.T) {
		constraint := parseConstraints("class: 1-3 or 5-7\nrow: 6-11 or 33-44")
		if len(constraint) != 4 {
			t.Fatalf("Wrong constraint parsed")
		}
		checkIn(t, Constraint{"class",1,3}, constraint)
		checkIn(t, Constraint{"class",5,7}, constraint)
		checkIn(t, Constraint{"row", 6,11}, constraint)
		checkIn(t, Constraint{"row", 33,44}, constraint)
	})
}

func TestParseMyTickets(t *testing.T) {
	t.Run("parseMyTickets should return correct result", func(t *testing.T) {
		myticket := parseMyTicket("your ticket:\n7,1,14")
		if !reflect.DeepEqual(myticket, []int{7,1,14}){
			t.Fatalf("Wrong ticket parsed for me")
		}
	})
}

func TestParseOtherTickets(t *testing.T) {
	t.Run("parseOtherTickets should return correct result", func(t *testing.T) {
		otherTickets := parseOtherTickets("nearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12")
		if len(otherTickets) != 4{
			t.Fatalf("Wrong ticket parsed for others")
		}
	})
}