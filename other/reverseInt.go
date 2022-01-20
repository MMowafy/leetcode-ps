package main

import (
	"fmt"
	"math/bits"
)

func main() {
	//fmt.Println(reverse(123))
	//fmt.Println(firstUniqChar("lleetcode"))
	fmt.Println(hammingWeight(32))
	fmt.Println(hammingWeight1(6565666))
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		rev = rev*10 + pop
	}
	return rev
}

func firstUniqChar(s string) int {
	strByte := []byte(s)
	mapStr := make(map[byte]int)
	for _, value := range strByte {
		mapStr[value] = mapStr[value] + 1
	}
	for index, value := range strByte {
		if mapStr[value] == 1 {
			return index
		}
	}
	return 0
}

func isAnagram(s string, t string) bool {

	return true
}

func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}

func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		fmt.Println(num)
		num = num & (num - 1)
		fmt.Println("esa", num)
		count++
	}
	return count
}
