package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carryOver := 0
	dummyNode := NewListNode(0)
	root := dummyNode
	for l1 != nil || l2 != nil {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}
		sum := val1 + val2 + carryOver
		carryOver = sum / 10
		root.Next = NewListNode(sum % 10)
		root = root.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carryOver > 0 {
		root.Next = NewListNode(carryOver)
	}
	return dummyNode.Next
}
