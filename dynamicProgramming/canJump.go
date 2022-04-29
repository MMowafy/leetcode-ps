package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	db := make([]int, len(nums))
	db[len(nums)-1] = 1
	return canJumpDB(nums, 0, db)
}

func canJumpDB(nums []int, position int, db []int) bool {
	fmt.Println(db)
	if (len(nums) - 1) == position {
		return true
	}
	if db[position] != 0 {
		if db[position] == 1 {
			return true
		}
		return false
	}
	furthestPos := int(math.Min(float64(position+nums[position]), float64(len(nums)-1)))
	fmt.Println("fur", furthestPos)
	for i := position + 1; i <= furthestPos; i++ {
		if canJumpDB(nums, i, db) {
			db[position] = 1
			return true
		}
	}
	db[position] = -1
	return false
}
