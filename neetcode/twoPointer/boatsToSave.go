package twoPointer

import (
	"fmt"
	"slices"
)

// limit 10
// 1, 2, 5, 5 = 6
//

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	res, l, r := 0, 0, len(people)-1

	for l <= r {
		remain := limit - people[r]
		r--
		res++
		if l <= r && remain >= people[l] {
			l++
		}
	}

	return res
}

// 1, 2, 4, 5, limit = 10

func TestNumRescueBoats() {
	tests := []struct {
		people []int
		limit  int
		exp    int
	}{
		{people: []int{1, 3, 2, 3, 2}, limit: 3, exp: 4},
		{people: []int{3, 2, 2, 1}, limit: 3, exp: 3},
		{people: []int{5, 1, 4, 2}, limit: 10, exp: 2},
	}

	for _, tt := range tests {
		got := numRescueBoats(tt.people, tt.limit)
		fmt.Printf("exp: %d, got: %d, passed: %t\n", tt.exp, got, tt.exp == got)
	}
}
