package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	sortedStrs := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		tempStr := []byte(strs[i])
		sort.Slice(tempStr, func(i, j int) bool {
			return tempStr[i] < tempStr[j]
		})
		sortedStrs[i] = string(tempStr)
	}
	var sortedKEYSIndex = make(map[string]int)
	var results [][]string
	for i := 0; i < len(strs); i++ {
		if _, ok := sortedKEYSIndex[sortedStrs[i]]; ok {
			index := sortedKEYSIndex[sortedStrs[i]]
			results[index] = append(results[index], strs[i])
		} else {
			results = append(results, []string{strs[i]})
			sortedKEYSIndex[sortedStrs[i]] = len(results) - 1
		}
	}
	return results
}
