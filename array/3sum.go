package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			rs := twoSum(nums, i)
			result = append(result, rs...)
		}
	}
	return result
}

func twoSum(nums []int, startIndex int) [][]int {
	var results [][]int
	seen := make(map[int]bool)
	target := -1 * nums[startIndex]
	for i := startIndex + 1; i < len(nums); i++ {
		remaining := target - nums[i]
		if _, ok := seen[remaining]; ok {
			results = append(results, []int{nums[startIndex], nums[i], remaining})
			for {
				if i+1 == len(nums) || nums[i] != nums[i+1] {
					break
				}
				i++
			}
		}
		seen[nums[i]] = true
	}
	return results
}
