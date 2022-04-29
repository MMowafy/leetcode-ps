package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	var result []string
	prev := lower - 1
	for i := 0; i <= len(nums); i++ {
		current := 0
		if i == len(nums) {
			current = upper + 1
		} else {
			current = nums[i]
		}
		if prev+1 == current-1 {
			result = append(result, strconv.Itoa(prev+1))
		}
		if prev+1 < current-1 {
			result = append(result, fmt.Sprintf("%d->%d", prev+1, current-1))
		}
		prev = current

	}
	return result
}
