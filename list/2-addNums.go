package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if getLen(l1) >= getLen(l2) {
		l1Head := l1
		carry := 0 //进位数
		for l1 != nil && l2 != nil {
			add := l1.Val + l2.Val + carry
			carry = add / 10
			l1.Val = add % 10

			if l1.Next == nil && carry != 0 { // [5] [5]特殊情况
				l1.Next = new(ListNode)
			}

			l1 = l1.Next
			l2 = l2.Next
		}

		for carry != 0 {
			add := l1.Val + carry
			carry = add / 10
			l1.Val = add % 10

			if l1.Next == nil && carry != 0 { // [9999] [99]特殊情况
				l1.Next = new(ListNode)
			}
			l1 = l1.Next
		}

		return l1Head
	} else {
		l2Head := l2
		carry := 0 //进位数
		for l1 != nil && l2 != nil {
			add := l1.Val + l2.Val + carry
			carry = add / 10
			l2.Val = add % 10

			if l2.Next == nil && carry != 0 {
				l2.Next = new(ListNode)
			}

			l1 = l1.Next
			l2 = l2.Next
		}

		for carry != 0 {
			add := l2.Val + carry
			carry = add / 10
			l2.Val = add % 10

			if l2.Next == nil && carry != 0 {
				l2.Next = new(ListNode)
			}
			l2 = l2.Next
		}

		return l2Head
	}

}

func getLen(l *ListNode) int {
	if l == nil {
		return 0
	}
	count := 1
	for l.Next != nil {
		count++
		l = l.Next
	}
	return count
}

func main() {
	// expect: 7 0 8
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}} // [2, 4, 3]
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}} // [5, 6, 4]

	result := addTwoNumbers(l1, l2)

	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}
