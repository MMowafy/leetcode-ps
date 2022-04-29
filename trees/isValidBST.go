package main

import (
	"fmt"
	helperPck "leetcode-problems/trees/helper"
)

func main() {
	l2 := &helperPck.TreeNode{
		Val: 0,
	}
	fmt.Println(isValidBST(l2))
}

func isValidBST(root *helperPck.TreeNode) bool {
	return isValidBSTDC(root)
}

var previousValue int = -1

func isValidBSTDC(root *helperPck.TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBSTDC(root.Left) {
		return false
	}
	if root.Val <= previousValue {
		return false
	}
	fmt.Println(root.Val)
	fmt.Println(previousValue)
	previousValue = root.Val
	fmt.Println(previousValue)
	return isValidBSTDC(root.Right)
}
