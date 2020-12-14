package main

import "testing"

func TestParseAddressAndValue(t *testing.T) {
	testCases := []struct {
		line string
		address int64
		value int64
	}{
		{"mem[8] = 11",8,11},
		{"mem[7] = 101",7,101},
	}
	for _, test := range testCases {
		t.Run("parseAddressAndValue should return correct result", func(t *testing.T) {
			gotAddress, gotValue := parseAddressAndValue(test.line)
			if gotAddress != test.address || gotValue != test.value {
				t.Fatalf("Wrong result expect (%d, %d), got (%d, %d)", test.address, test.value, gotAddress, gotValue)
			}
		})
	}
}

func TestUpdateValue(t *testing.T) {
	testCases := []struct {
		mask string
		value int64
		updated int64
	}{
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",11,73},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",101,101},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",0,64},
	}
	for _, test := range testCases {
		t.Run("updateValue should return correct result", func(t *testing.T) {
			got := updateValue(test.value, test.mask)
			if got != test.updated {
				t.Fatalf("Wrong result expect %d, got %d", test.updated, got)
			}
		})
	}
}