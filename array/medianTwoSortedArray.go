package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result []int
	j := 0
	i := 0
	for k := 0; k < len(nums1)+len(nums2); k++ {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				result = append(result, nums1[i])
				i++
				continue
			} else {

				result = append(result, nums2[j])
				j++
				continue
			}
		}

		if i < len(nums1) {
			result = append(result, nums1[i:]...)
			break
		}

		if j < len(nums2) {
			result = append(result, nums2[j:]...)
			break
		}

	}
	return medianOf(result)
}

func medianOf(nums []int) float64 {
	l := len(nums)

	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}

	return float64(nums[l/2])
}
