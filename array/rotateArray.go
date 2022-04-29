package main

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/
func rotate(nums []int, k int) {
	results := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tempIndex := (i + k) % len(nums)
		results[tempIndex] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = results[i]
	}
}
