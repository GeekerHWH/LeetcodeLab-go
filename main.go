package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if getLen(l1) >= getLen(l2) {
		l1Head := l1
		carry := 0 //进位数
		for l1.Next != nil && l2.Next != nil {
			add := l1.Val + l2.Val + carry
			carry = add / 10
			l1.Val = add % 10

			l1 = l1.Next
			l2 = l2.Next
		}

		for carry != 0 {
			add := l1.Val + carry
			carry = add / 10
			l1.Val = add % 10

			l1 = l1.Next
		}

		return l1Head
	} else {
		l2Head := l2
		carry := 0 //进位数
		for l1.Next != nil && l2.Next != nil {
			add := l1.Val + l2.Val + carry
			carry = add / 10
			l2.Val = add % 10

			l1 = l1.Next
			l2 = l2.Next
		}

		for carry != 0 {
			add := l2.Val + carry
			carry = add / 10
			l2.Val = add % 10

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

}
