package main

import helperPck "leetcode-problems/trees/helper"

func main() {

}

func sortedArrayToBST(nums []int) *helperPck.TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) *helperPck.TreeNode {
	if left > right {
		return nil
	}
	middle := left + right/2
	if (left+right/2)%2 == 1 {
		middle = middle + 1
	}
	node := helperPck.TreeNode{
		Val: nums[middle],
	}
	node.Left = helper(nums, left, middle)
	node.Right = helper(nums, middle+1, right)
	return &node
}
