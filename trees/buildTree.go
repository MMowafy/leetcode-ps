package main

import (
	"fmt"
	helperPck "leetcode-problems/trees/helper"
)

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

var preOrderIndex = 0

func buildTree(preorder []int, inorder []int) *helperPck.TreeNode {
	inorderIndexMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inorderIndexMap[inorder[i]] = i
	}
	preOrderIndex = 0
	return buildTreeHelper(preorder, 0, len(preorder)-1, inorderIndexMap)
}

func buildTreeHelper(preorder []int, left int, right int, inorderIndexMap map[int]int) *helperPck.TreeNode {
	if left > right {
		return nil
	}
	root := &helperPck.TreeNode{
		Val: preorder[preOrderIndex],
	}
	preOrderIndex++
	root.Left = buildTreeHelper(preorder, left, inorderIndexMap[root.Val]-1, inorderIndexMap)
	root.Right = buildTreeHelper(preorder, inorderIndexMap[root.Val]+1, right, inorderIndexMap)
	return root
}
