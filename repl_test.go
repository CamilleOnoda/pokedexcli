package main

import "testing"

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "extra space at the end",
			input:    " Hello World ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "extra spaces at the beginning",
			input:    "   Hello World",
			expected: []string{"hello", "world"},
		},
		{
			name:     "multiple spaces at the beginning and end",
			input:    "I     Love Go    ",
			expected: []string{"i", "love", "go"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only spaces",
			input:    "      ",
			expected: []string{},
		},
		{
			name:     "mixed case",
			input:    "GoLanG Is FuN",
			expected: []string{"golang", "is", "fun"},
		},
	}

	for _, test := range tests {
		actual := cleanInput(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("Input: %v | Expected: %v | Got: %v", test.input, test.expected, actual)
		}
		for i := range actual {
			if actual[i] != test.expected[i] {
				t.Errorf("Expected: %v | Got: %v", test.expected[i], actual[i])
			}
		}

	}
}
