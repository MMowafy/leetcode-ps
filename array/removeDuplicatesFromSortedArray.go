package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	insertIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
	}
	fmt.Println(nums)
	return insertIndex
}
