package main

/*
1. imagine the usual scenario
2. identify general edge cases
3. identify edge cases specific on the implementation
*/

import (
	"testing"
)

func TestSolution(t *testing.T) {
	for i, tc := range []struct{
		S string
		expected int
	}{
		// their examples.
		{"abbabba", 4},
		{"codility", 0},
	}{
		actual := Solution(tc.S)
		if actual != tc.expected {
			t.Errorf("case %d: (%v) expected %d, got %d", i, tc.S, tc.expected, actual)
		}
	}
}


