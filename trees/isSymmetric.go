package main

import helperPck "leetcode-problems/trees/helper"

func main() {

}

func isSymmetric(root *helperPck.TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(root1 *helperPck.TreeNode, root2 *helperPck.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && isMirror(root1.Right, root2.Left) && isMirror(root1.Left, root2.Right)
}
