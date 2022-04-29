package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestPalindrome("ac"))
}

var longestCommon string

func longestPalindrome(s string) string {
	longestCommon = ""
	strByte := []byte(s)
	lenStr := len(strByte)
	dp := make([][]int, lenStr)
	for i := 0; i < lenStr; i++ {
		dp[i] = make([]int, lenStr) // initialize a slice of dx unit8 in each of dy slices
	}
	LCSubstring(dp, strByte, 0, lenStr-1)
	return longestCommon
}

func LCSubstring(db [][]int, s1 []byte, startIndex int, endIndex int) []byte {
	if startIndex > endIndex {
		return nil
	}
	if startIndex == endIndex {
		return []byte{s1[startIndex]}
	}
	var str1 []byte
	if db[startIndex][endIndex] == 0 {
		if s1[startIndex] == s1[endIndex] {
			remainingString := LCSubstring(db, s1, startIndex+1, endIndex-1)
			if len(remainingString) == endIndex-startIndex-1 {
				str1 = s1[startIndex:endIndex]
			}
		}

		str2 := LCSubstring(db, s1, startIndex+1, endIndex)
		str3 := LCSubstring(db, s1, startIndex, endIndex-1)
		max := int(math.Max(math.Max(float64(len(str1)), float64(len(str2))), float64(len(str3))))
		if max == int(lcs) {
			longestCommon = string(s1[startIndex:endIndex])
		} else if max == int(lcs2) {
			longestCommon = string(s1[startIndex+1 : endIndex])
		} else {
			longestCommon = string(s1[startIndex : endIndex-1])
		}
		db[startIndex][endIndex] = int(math.Max(math.Max(lcs, lcs2), lcs3))
	}
	return db[startIndex][endIndex]
}
