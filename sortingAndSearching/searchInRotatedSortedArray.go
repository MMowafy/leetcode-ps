package main

import "fmt"

func main() {
	fmt.Println(search([]int{1}, 1))
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return nums[0]
		}
		return -1
	}

	rotationIndex := findRotationIndex(nums)
	fmt.Println(rotationIndex)

	if nums[rotationIndex] == target {
		return rotationIndex
	}
	if rotationIndex == 0 {
		return searchDC(nums, target, 0, len(nums)-1)
	}
	if target < nums[0] {
		return searchDC(nums, target, rotationIndex+1, len(nums)-1)
	}
	return searchDC(nums, target, 0, rotationIndex)

}

func searchDC(nums []int, target int, left int, right int) int {
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func findRotationIndex(nums []int) int {
	//find rotationIndex
	left := 0
	right := len(nums) - 1
	rotationIndex := 0
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] > nums[middle+1] {
			rotationIndex = middle + 1
			break
		} else {
			if nums[middle] < nums[left] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return rotationIndex
}
