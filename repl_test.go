package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Quaxly is the greatest pokemon",
			expected: []string{"quaxly", "is", "the", "greatest", "pokemon"},
		},
		{
			input:    "ro foe Lo 892 hello?xx",
			expected: []string{"ro", "foe", "lo", "892", "hello?xx"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("error: length of %s does not match", actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("error: %s does not match expected word: %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
