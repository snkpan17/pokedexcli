package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []Case{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " PIKACHU  Salamander raiCHU",
			expected: []string{"pikachu", "salamander", "raichu"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, test := range cases {
		actual := cleanInput(test.input)
		expected := test.expected

		if len(actual) != len(expected) {
			fmt.Printf("actual: %v, expected: %v", actual, expected)
			t.Fatalf("Length of slice is different: actual len - %d, expected len - %d", len(actual), len(expected))
		}
		for i := range actual {
			accStr := actual[i]
			expStr := expected[i]
			if accStr != expStr {
				fmt.Printf("actual: %v, expected: %v", actual, expected)
				t.Fatalf("String don't match: actual %s, expected %s", accStr, expStr)
			}
		}
	}
}
