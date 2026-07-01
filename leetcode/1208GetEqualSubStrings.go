package leetcode

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
	result := 0
	left := 0
	curCost := 0
	for right := range len(s) {
		curCost += abs(int(s[right]) - int(t[right]))
		for curCost > maxCost {
			curCost -= abs(int(s[left]) - int(t[left]))
			left++
		}

		result = max(result, right-left+1)
	}

	return result
}

func TestEqualSubstring() {
	tests := []struct {
		s       string
		t       string
		maxCost int
		exp     int
	}{
		{"abcd", "bcdf", 3, 3},
	}

	for _, tt := range tests {
		got := equalSubstring(tt.s, tt.t, tt.maxCost)
		fmt.Printf("exp: %d, got: %d\n", tt.exp, got)
	}
}
