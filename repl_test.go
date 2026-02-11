package main

import (
	"testing"
)
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input: "     some    test123",
			expected: []string{"some", "test123"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]

			if word != expected {
				t.Errorf("error: %v", actual)
			}
		}
	}	
}