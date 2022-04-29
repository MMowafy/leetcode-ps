package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	return coinChangeDP(coins, amount, dp)
}

func coinChangeDP(coins []int, amount int, dp []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if dp[amount] != 0 {
		return dp[amount]
	}

	min := math.MaxInt
	for _, coin := range coins {
		fmt.Println("coin ", coin)
		fmt.Println("amount ", amount)
		res := coinChangeDP(coins, amount-coin, dp)
		fmt.Println("ress ", res)
		if res >= 0 && res < min {
			min = 1 + res
		}
		fmt.Println("min", min)
		fmt.Println()
	}

	dp[amount] = min
	if min == math.MaxInt {
		dp[amount] = -1
	}
	return dp[amount]
}
