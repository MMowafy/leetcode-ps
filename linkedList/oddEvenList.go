package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	size := 0
	for head != nil {
		head = head.Next
		size++
	}
	for i := 0; i < size/2; size++ {
		tempNext := head.Next
		head.Next = head.Next.Next
		head.Next.Next = tempNext
		tempNext.Next = tempNext.Next.Next
		head = head.Next
	}
	return head
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
