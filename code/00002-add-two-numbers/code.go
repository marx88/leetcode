package main

import (
	"log"
)

// ListNode 结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				9,
				nil,
			},
		},
	}
	l2 := &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				&ListNode{
					9,
					nil,
				},
			},
		},
	}
	t := addTwoNumbers(l1, l2)
	for t != nil {
		log.Println(*t)
		t = t.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{l1.Val + l2.Val, nil}
	temp := first
	for l1.Next != nil || l2.Next != nil || temp.Val > 9 {
		temp.Next = &ListNode{temp.Val / 10, nil}
		temp.Val = temp.Val % 10
		temp = temp.Next
		if l1.Next != nil {
			temp.Val += l1.Next.Val
			l1 = l1.Next
		}
		if l2.Next != nil {
			temp.Val += l2.Next.Val
			l2 = l2.Next
		}
	}
	return first
}
