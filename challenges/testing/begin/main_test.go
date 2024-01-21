// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"a2 32 ^ &/)", 1},
		{"812 %6ab//", 2},
	}
	counter := letterCounter{identifier: "letters"}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if actual := counter.count(tc.input); actual != tc.expected {
				t.Errorf("actual %v != want %v", actual, tc.expected)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"abc 1,?/", 1},
		{"abc 12&8 ^", 3},
	}
	counter := numberCounter{designation: "numbers"}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if actual := counter.count(tc.input); actual != tc.expected {
				t.Errorf("actual %v != want %v", actual, tc.expected)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"abc 1,?/", 4},
		{"abc 12&8 ^_", 5},
	}
	counter := symbolCounter{label: "symbols"}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if actual := counter.count(tc.input); actual != tc.expected {
				t.Errorf("actual %v != want %v", actual, tc.expected)
			}
		})
	}
}
