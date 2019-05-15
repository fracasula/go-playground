package main

import (
	"github.com/davecgh/go-spew/spew"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val := 0
	if l1 != nil {
		val += l1.Val
	}
	if l2 != nil {
		val += l2.Val
	}

	if val > 9 {
		val = val - 10
		if l1 != nil {
			if l1.Next == nil {
				l1.Next = &ListNode{Val: 1}
			} else {
				l1.Next.Val += 1
			}
		}
	}
	node := &ListNode{Val: val}
	var n1, n2 *ListNode
	if l1 != nil {
		n1 = l1.Next
	}
	if l2 != nil {
		n2 = l2.Next
	}
	if n1 != nil || n2 != nil {
		node.Next = addTwoNumbers(n1, n2)
	}

	return node
}

type scenario struct {
	l1, l2   *ListNode
	expected string
}

func main() {
	scenarios := []scenario{
		{
			l1:       &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:       &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expected: "<*>{7 <*>{0 <*>{8 <nil>}}}",
			// 7, 0, 8
		},
		{
			l1:       &ListNode{Val: 5},
			l2:       &ListNode{Val: 5},
			expected: "<*>{0 <*>{1 <nil>}}",
			// 0, 1
		},
		{
			l1:       &ListNode{Val: 9},
			l2:       &ListNode{Val: 9},
			expected: "<*>{8 <*>{1 <nil>}}",
			// 8, 1
		},
		{
			l1:       &ListNode{Val: 1, Next: &ListNode{Val: 8}},
			l2:       &ListNode{Val: 0},
			expected: "<*>{1 <*>{8 <nil>}}",
			// 1, 8
		},
	}

	for i, s := range scenarios {
		actual := spew.Sprint(addTwoNumbers(s.l1, s.l2))
		if s.expected != actual {
			log.Fatalf(
				"Scenario %d failed, got %q instead of %q", i+1, actual, s.expected,
			)
		}
	}
}
