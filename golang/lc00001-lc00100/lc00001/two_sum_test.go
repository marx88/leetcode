package lc00001

import (
	"testing"
)

type caseType struct {
	nums     []int
	target   int
	expected []int
}

var cases = []caseType{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
}

func TestTwoSum(t *testing.T) {
	for _, c := range cases {
		result := twoSum(c.nums, c.target)
		if result[0] != c.expected[0] || result[1] != c.expected[1] {
			t.Fatalf("twoSum function failed, nums: %v, target: %d, execpted: %v, result: %v", c.nums, c.target, c.expected, result)
		}
	}
}

func TestTwoSumMap(t *testing.T) {
	for _, c := range cases {
		result := twoSumMap(c.nums, c.target)
		if result[0] != c.expected[0] || result[1] != c.expected[1] {
			t.Fatalf("twoSum function failed, nums: %v, target: %d, execpted: %v, result: %v", c.nums, c.target, c.expected, result)
		}
	}
}
