package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5, 6, 7}))
}

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/
func maxProfit(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}
