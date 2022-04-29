package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 1, 3}))
}

func increasingTriplet(nums []int) bool {
	i := 0
	min := math.MaxInt
	secondMin := math.MaxInt
	for i < len(nums) {
		if nums[i] <= min {
			min = nums[i]
		} else if nums[i] <= secondMin {
			secondMin = nums[i]
		} else {
			return true
		}
		fmt.Println(min)
		fmt.Println(secondMin)
		i++
	}

	return false
}
