package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	l1 := &Node{
		Val: 7,
	}
	l2 := &Node{
		Val: 6,
	}
	l3 := &Node{
		Val: 5,
	}
	l4 := &Node{
		Val: 4,
	}
	l5 := &Node{
		Val:   3,
		Left:  l2,
		Right: l1,
	}
	l6 := &Node{
		Val:   2,
		Left:  l4,
		Right: l3,
	}
	l7 := &Node{
		Val:   1,
		Left:  l6,
		Right: l5,
	}
	root2 := connect(l7)
	LevelOrderTraversal22(root2)
}

func connect(root *Node) *Node {
	return levelOrderTraversal(nil)
}

func levelOrderTraversal(root *Node) *Node {
	if root == nil {
		return nil
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		queueLength := len(queue)
		fmt.Println(queueLength)
		for i := 0; i < queueLength; i++ {
			popped := queue[0]
			queue = queue[1:]
			if i != queueLength-1 {
				popped.Next = queue[0]
			}
			if popped.Left != nil {
				queue = append(queue, popped.Left)
			}
			if popped.Right != nil {
				queue = append(queue, popped.Right)
			}
		}
	}
	return root
}

func LevelOrderTraversal22(root *Node) {
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		fmt.Println()
		fmt.Println("node vale", node.Val)
		fmt.Println("next", node.Next)
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
