package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	strRune := []rune(pattern)
	mapStr := make(map[string]string)
	mapStr2 := make(map[string]string)
	if len(strs) != len(strRune) {
		return false
	}
	for index, val := range strRune {
		char := string(val)
		_, ok := mapStr2[strs[index]]
		if ok {
			if mapStr2[strs[index]] != char {
				return false
			}
		} else {
			mapStr2[strs[index]] = char
		}

		_, ok = mapStr[char]
		if ok {
			if mapStr[char] != strs[index] {
				return false
			}
		} else {
			mapStr[char] = strs[index]
		}
	}

	return true
}
