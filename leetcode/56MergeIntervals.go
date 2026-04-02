package leetcode

import (
	"slices"
)

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	result := [][]int{}
	cur := intervals[0]

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]

		if next[0] <= cur[1] {
			if next[1] > cur[1] {
				cur[1] = next[1]
			}
		} else {
			result = append(result, cur)
			cur = next
		}
	}

	return result
}
