package main

import (
	"fmt"
	helper2 "leetcode-problems/trees/helper"
)

func main() {
	l1 := &helper2.TreeNode{
		Val: 1,
	}
	l2 := &helper2.TreeNode{
		Val: 3,
	}
	l3 := &helper2.TreeNode{
		Val:   2,
		Left:  l1,
		Right: l2,
	}
	fmt.Println(inorderSuccessor(l3, l1))
}

//func inorderSuccessor(root *helper2.TreeNode, p *helper2.TreeNode) *helper2.TreeNode {
//	var successor *helper2.TreeNode
//	for root != nil {
//		if p.Val >= root.Val {
//			root = root.Right
//		} else {
//			successor = root
//			root = root.Left
//		}
//	}
//	return successor
//}
var previous *TreeNode
var inordersuccessor *TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	previous = nil
	inordersuccessor = nil
	if p.Right != nil {
		root = root.Right
		leftMost := root
		fmt.Println(leftMost)
		for root.Left != nil {
			leftMost = root.Left
			root = root.Left
		}
		return leftMost
	} else {
		inorderSuccessorDC(root, p)
		return inordersuccessor
	}
}

func inorderSuccessorDC(root *TreeNode, p *TreeNode) {
	if root == nil {
		return
	}
	inorderSuccessorDC(root.Left, p)
	if previous == p {
		inordersuccessor = root
	}
	previous = root
	inorderSuccessorDC(root.Right, p)
}
