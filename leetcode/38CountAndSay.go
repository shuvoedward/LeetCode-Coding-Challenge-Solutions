package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// if one frequency - count it and add to the front
// if more than one frequency - count and replace the front
// 21 = 2 unique numbers = 2 appears 1 time, and 1 appears 1 time.
// add each count to its front = 1211
// 1211 = 111221
// keep left to where needs to be replaced or inserted
// right for counting, substring, conditions -> count frequency,
// if one -> add to left position
// if more than one -> insert to left position, then compress the rest of it.

func countAndSay(n int) string {
	var result = "1"
	if n == 1 {
		return result
	}

	var curStr strings.Builder
	l := 0
	r := 0
	count := 0

	for i := 2; i <= n; {
		count = 0
		l = 0
		for r < len(result) {
			if result[l] != result[r] {
				curStr.WriteString(strconv.Itoa(count))
				curStr.WriteByte(result[l])
				l = r
				count = 0
			} else {
				count++
				r++
			}
		}
		curStr.WriteString(strconv.Itoa(count))
		curStr.WriteByte(result[l])
		result = curStr.String()
		curStr.Reset()
		i++
		r = 0
	}

	return result
}

func RunCountAndStay() {
	tests := []struct {
		n   int
		exp string
	}{
		// {n: 1, exp: "1"},
		{n: 3, exp: "21"},
		{n: 4, exp: "1211"},
		{n: 10, exp: "1211"},
	}

	for _, tt := range tests {
		got := countAndSay(tt.n)
		fmt.Printf("exp: %s, got: %s\n", tt.exp, got)
	}
}
