package main

import "testing"

func testingCleanInput(t *testing.T) {
	cs := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello World ",
			expected: []string{"Hello", "World"},
		},
		{
			input:    "   Hello World",
			expected: []string{"Hello", "World"},
		},
		{
			input:    "I     Love Go    ",
			expected: []string{"I", "love", "Go"},
		},
	}

	for _, c := range cs {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Input: %v | Expected: %v | Got: %v", c.input, c.expected, actual)
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Expected: %v | Got: %v", c.expected[i], actual[i])
			}
		}

	}
}
