package lc00003

import "testing"

type caseType struct {
	str      string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []caseType{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, c := range cases {
		result := lengthOfLongestSubstring(c.str)
		if result != c.expected {
			t.Fatalf("lengthOfLongestSubstring function failed, str: %s, expected: %d, result: %d", c.str, c.expected, result)
		}
	}
}
