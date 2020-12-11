package lc00004

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	type caseType struct {
		nums1    []int
		nums2    []int
		excepted float64
	}

	cases := []caseType{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{1}, 1},
		{[]int{2}, []int{}, 2},
	}

	for _, c := range cases {
		result := findMedianSortedArrays(c.nums1, c.nums2)
		if result != c.excepted {
			t.Fatalf("findMedianSortedArrays function failed, nums1: %v, nums2: %v, execpted: %f, result: %f", c.nums1, c.nums2, c.excepted, result)
		}
	}
}
