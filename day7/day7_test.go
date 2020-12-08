package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("Parse should parse bag contains no other bag", func(t *testing.T) {
		a, b := parse("dotted black bags contain no other bags.")
		if a != "dotted black" {
			t.Fatal("Wrong parsing for outer bag")
		}
		if len(b) != 0 {
			t.Fatal("Wrong parsing for inner bag")
		}
	})

	t.Run("Parse should parse bag contains some other bags", func(t *testing.T) {
		a, b := parse("muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.")
		if a != "muted yellow" {
			t.Fatal("Wrong parsing for outer bag")
		}
		if len(b) != 2 {
			t.Fatal("Wrong parsing for inner bag")
		}
		if b[0] != "shiny gold" && b[0] != "faded blue" {
			t.Fatal("Wrong parsing for inner bag")
		}
		if b[1] != "shiny gold" && b[1] != "faded blue" {
			t.Fatal("Wrong parsing for inner bag")
		}
	})

	t.Run("Parse should parse bag contains another bag", func(t *testing.T) {
		a, b := parse("muted yellow bags contain 2 shiny gold bags.")
		if a != "muted yellow" {
			t.Fatal("Wrong parsing for outer bag")
		}
		if len(b) != 1 {
			t.Fatal("Wrong parsing for inner bag")
		}
		if b[0] != "shiny gold" {
			t.Fatal("Wrong parsing for inner bag")
		}
	})
}


func TestParseWithCount(t *testing.T) {
	t.Run("parse_with_count() should parse bag contains no other bag", func(t *testing.T) {
		a, b := parse_with_count("dotted black bags contain no other bags.")
		if a != "dotted black" {
			t.Fatal("Wrong parsing for outer bag")
		}
		if len(b) != 0 {
			t.Fatal("Wrong parsing for inner bag")
		}
	})

	t.Run("parse_with_count() should parse bag contains some other bags", func(t *testing.T) {
		a, b := parse_with_count("muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.")
		if a != "muted yellow" {
			t.Fatal("Wrong parsing for outer bag")
		}
		if len(b) != 2 {
			t.Fatal("Wrong parsing for inner bag")
		}
		if b["shiny gold"] != 2 {
			t.Fatal("Wrong parsing for inner bag")
		}
		if b["faded blue"] != 9 {
			t.Fatal("Wrong parsing for inner bag")
		}
	})

	t.Run("parse_with_count() should parse bag contains another bag", func(t *testing.T) {
		a, b := parse_with_count("muted yellow bags contain 2 shiny gold bags.")
		if a != "muted yellow" {
			t.Fatal("Wrong parsing for outer bag")
		}
		if len(b) != 1 {
			t.Fatal("Wrong parsing for inner bag")
		}
		if b["shiny gold"] != 2 {
			t.Fatal("Wrong parsing for inner bag")
		}
	})
}