package main

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/
func containsDuplicate(nums []int) bool {
	mapList := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		_, ok := mapList[nums[i]]
		if ok {
			return true
		}
		mapList[nums[i]] = true
	}
	return false
}
