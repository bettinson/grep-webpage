package main

import (
	"strings"
	"testing"
)

func TestSearchReader(t *testing.T) {
	r := strings.NewReader("hey\nhey\njake\nyou're\nan\nall\nstar\nstar\n")
	lines := searchReader(r, func(s string) bool {
		return s == "hey"
	})
	if len(lines) != 2 {
		t.Errorf("expected two heys: %v", lines)
	}
}

func TestRegexPredicate(t *testing.T) {
	pattern := "^a*b*c$"
	matchString := "aabbc"
	matchFail := "aabb"

	f := regexPredicate(pattern)
	if !f(matchString) {
		t.Errorf("Pattern does not match")
	}
	if f(matchFail) {
		t.Errorf("Pattern should not match")
	}
}
