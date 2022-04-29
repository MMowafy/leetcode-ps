package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 2, 4}))
}

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
func singleNumber(nums []int) int {
	mapList := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		mapList[nums[i]] += 1
	}
	for i := 0; i < len(nums); i++ {
		val, _ := mapList[nums[i]]
		if val == 1 {
			return nums[i]
		}
	}
	return 0
}
