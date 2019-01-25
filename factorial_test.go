package main

import (
	"testing"
)

func TestFact(t *testing.T) {
	if fact(3) != 6 {
		t.Error("Expected 6")
	}
}

func TestTable(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{3, 6},
		{4, 24},
		{5, 120},
	}
	for _, test := range tests {
		output := fact(test.input)
		if output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, {} received ", test.input, test.expected, output)
		} else {
			t.Log("Test Passed: {} inputted, {} expected, {} received ", test.input, test.expected, output)
		}
	}
}
