package main

import (
	"strings"
	"testing"
)

func TestSearchReader(t *testing.T) {
	r := strings.NewReader("hey\nhey\njake\nyou're\nan\nall\nstar\n")
	lines := searchReader(r, func(s string) bool {
		return s == "hey\n"
	})
	if len(lines) != 2 {
		t.Errorf("expected two heys: %v", lines)
	}

}
