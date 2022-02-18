package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	kk := l3
	mid := 0
	for l1 != nil || l2 != nil || mid > 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += mid
		kk.Next = &ListNode{Val: sum % 10}
		kk = kk.Next
		mid = sum / 10
	}
	return l3.Next
}
