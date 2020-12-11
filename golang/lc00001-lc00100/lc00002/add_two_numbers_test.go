package lc00002

import "testing"

type ln ListNode

type caseType struct {
	l1       []int
	l2       []int
	expected int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []caseType{
		{[]int{0}, []int{0}, 0},
		{[]int{0}, []int{1}, 1},
		{[]int{1}, []int{0}, 1},
		{[]int{1, 0}, []int{1, 0, 2}, 112},
		{[]int{9, 4, 2}, []int{9, 4, 6, 5}, 10407},
	}

	for _, c := range cases {
		result := ln2Int(addTwoNumbers(int2Ln(c.l1), int2Ln(c.l2)))
		if result != c.expected {
			t.Fatalf("addTwoNumbers function failed, l1: %v, l2: %v, expected: %d, result: %d", c.l1, c.l2, c.expected, result)
		}
	}
}

func int2Ln(num []int) *ListNode {
	var tn *ListNode
	for i, length := 0, len(num); i < length; i++ {
		tn = &ListNode{num[i], tn}
	}
	return tn
}

func ln2Int(l *ListNode) int {
	num := 0
	for i := 1; l != nil; i *= 10 {
		num += l.Val * i
		l = l.Next
	}
	return num
}
