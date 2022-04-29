package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 1, Next: nil}
	l2 := &ListNode{Val: 2, Next: l1}
	l3 := &ListNode{Val: 2, Next: l2}
	l4 := &ListNode{Val: 1, Next: l3}
	fmt.Println(isPalindrome(l4))
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func isPalindrome(head *ListNode) bool {
	reversedList := reverseList(head)
	for {
		if head == nil {
			return true
		}
		fmt.Println("head ", head)
		fmt.Println("reversed ", reversedList)
		if head.Val != reversedList.Val {
			return false
		}
		head = head.Next
		reversedList = reversedList.Next
	}
	return true
}

func NewStack() *Stack {
	return &Stack{
		List: linked_list.NewLinkedList(),
	}
}

func (stack *Stack) Push(val int) {
	fmt.Println("Push ", val)
	stack.List.InsertAt(val, 0)
}

func (stack *Stack) Pop() {
	fmt.Println("Pop ", stack.List.Head.Value)
	stack.List.DeleteNodeAt(0)
}
