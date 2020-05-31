package main

import "testing"

func TestFindWord(t *testing.T) {
	type test struct {
		id       int
		dice     int
		expected string
	}

	tests := []test{
		{2345, 4, "diner"},
		{5126, 4, "riot"},
		{22151, 5, "danger"},
		{25141, 5, "email"},
	}

	for _, v := range tests {
		actual := findWord(v.id, v.dice)
		if actual != v.expected {
			t.Error("Expected:", v.expected, "- Got:", actual)
		}
	}
}
