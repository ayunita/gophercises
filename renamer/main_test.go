package main

import "testing"

func TestPretty(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"test_001.txt", "Test 001.txt"},
		{"good_morning_1.txt", "Good_morning 1.txt"},
		{"Hello world_004.txt", "Hello World 004.txt"},
		{"12931_004.txt", "12931 004.txt"},
		{"good morning (1 of 10).txt", "Good Morning 1.txt"},
		{"good_morning_(5 of 10).txt", "Good_morning_ 5.txt"},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if got := Pretty(tc.input); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}
