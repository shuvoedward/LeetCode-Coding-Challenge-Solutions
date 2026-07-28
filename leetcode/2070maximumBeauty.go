package leetcode

import (
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {

	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	n := len(items)
	maxBeauty := make([]int, n)
	prices := make([]int, n)
	best := 0
	for i := 0; i < n; i++ {
		if items[i][1] > best {
			best = items[i][1]
		}
		maxBeauty[i] = best
		prices[i] = items[i][0]
	}

	ans := make([]int, len(queries))
	for idx, q := range queries {
		// First index where prices[i] > q
		pos := sort.Search(n, func(i int) bool {
			return prices[i] > q
		})
		if pos == 0 {
			ans[idx] = 0
		} else {
			ans[idx] = maxBeauty[pos-1]
		}
	}

	return ans
}
