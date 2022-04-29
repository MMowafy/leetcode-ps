package main

import (
	"math"
	"sort"
)

func main() {

}

func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastInternalEnd := res[len(res)-1][1]
		if intervals[i][0] <= lastInternalEnd {
			res[len(res)-1][1] = int(math.Max(float64(lastInternalEnd), float64(intervals[i][1])))
		} else {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return res
}
