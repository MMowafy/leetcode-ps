package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePathsBottomUp(3, 7))
}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	rightWaysCount := uniquePaths(m-1, n)
	downWaysCount := uniquePaths(m, n-1)
	return rightWaysCount + downWaysCount
}

func uniquePathsBottomUp(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for coloumn := 1; coloumn < n; coloumn++ {
		for row := 1; row < m; row++ {
			dp[row][coloumn] = dp[row-1][coloumn] + dp[row][coloumn-1]
		}
	}
	return dp[m-1][n-1]
}
