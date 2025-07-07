package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},

		{
			input:    "",
			expected: []string{},
		},

		{
			input:    "hello",
			expected: []string{"hello"},
		},
	}

	for i, c := range cases {
		isCorrect := true

		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			isCorrect = false
		} else if len(actual) != 0 && len(c.expected) != 0 {

			for j := range actual {
				if actual[j] != c.expected[j] {
					isCorrect = false
				}
			}

		}

		if !isCorrect {
			t.Errorf("Test #%v \nIncorrect Value: %v\nExpected: %v\n", i+1, actual, c.expected)
		} else {
			fmt.Println("Test #", i+1, "\nCorrect Value:", actual, "\nExpected:", c.expected, "\n ")
		}

	}

}
