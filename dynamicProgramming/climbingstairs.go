package main

import "fmt"

func main() {
	fmt.Println(climbStairs(45))
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	return helper(n, dp)
}

func helper(n int, dp []int) int {
	if dp[n] > 0 {
		return dp[n]
	}
	if n <= 1 {
		return 1
	}
	numberofWays1 := helper(n-1, dp)
	numberofWays2 := helper(n-2, dp)
	dp[n] = numberofWays2 + numberofWays1
	return dp[n]
}
