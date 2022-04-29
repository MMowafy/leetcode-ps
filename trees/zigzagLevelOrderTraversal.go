package main

import (
	"fmt"
	helperPck "leetcode-problems/trees/helper"
)

func main() {
	l1 := &helperPck.TreeNode{
		Val: 7,
	}
	l2 := &helperPck.TreeNode{
		Val: 15,
	}
	l3 := &helperPck.TreeNode{
		Val:   20,
		Left:  l2,
		Right: l1,
	}
	l4 := &helperPck.TreeNode{
		Val: 9,
	}
	l5 := &helperPck.TreeNode{
		Val:   3,
		Left:  l4,
		Right: l3,
	}
	fmt.Println(zigzagLevelOrder(l5))
}

func zigzagLevelOrder(root *helperPck.TreeNode) [][]int {
	var queue []*helperPck.TreeNode
	queue = append(queue, root)
	var levelsArr [][]int
	level := 0
	isLeftOrder := true
	for len(queue) > 0 {
		fmt.Println("HELLLLLOOOOOOOOOOOOOOOOO")
		queueLength := len(queue)
		var tempInnerArr []int
		for i := 0; i < queueLength; i++ {
			popped := queue[0]
			fmt.Println("popped", popped)
			tempInnerArr = append(tempInnerArr, popped.Val)
			queue = queue[1:]
			if isLeftOrder {
				fmt.Println("ifff ")
				if popped.Left != nil {
					queue = append(queue, popped.Left)
				}
				if popped.Right != nil {
					queue = append(queue, popped.Right)
				}
			} else {
				fmt.Println("hello from else")
				if popped.Right != nil {
					queue = append(queue, popped.Right)
				}
				if popped.Left != nil {
					queue = append(queue, popped.Left)
				}
			}

		}
		fmt.Println("last", tempInnerArr)
		if len(levelsArr) != 0 {
			isLeftOrder = !isLeftOrder
		}
		levelsArr = append(levelsArr, tempInnerArr)
		level++
	}
	return levelsArr
}
