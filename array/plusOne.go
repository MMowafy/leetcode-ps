package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 1, 2, 4}))
}

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/559/
func plusOne(digits []int) []int {
	carryOver := 1
	for i := (len(digits) - 1); i >= 0; i-- {
		digits[i] = digits[i] + carryOver
		if digits[i] > 9 {
			carryOver = 1
			digits[i] = 0
		} else {
			carryOver = 0
			break
		}
	}
	a := digits
	if carryOver == 1 {
		a = []int{1}
		a = append(a, digits...)
	}
	return a
}
