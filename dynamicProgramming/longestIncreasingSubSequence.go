package main

import "math"

func main() {

}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(nums[i]), float64(dp[j]+1)))
			}
		}
	}
	longest := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > longest {
			longest = dp[i]
		}
	}
	return longest
}
