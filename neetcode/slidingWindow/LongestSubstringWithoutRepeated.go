package slidingwindow

import "fmt"

func longestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	seen := map[rune]int{}

	l := 0
	maxLen := 0

	for r, ch := range s {
		if index, exists := seen[ch]; exists && l <= index {
			l = index + 1
		}

		seen[ch] = r

		if r-l+1 > maxLen {
			fmt.Println(r)
			maxLen = r - l + 1
		}
	}

	return maxLen
}

func TestLongestSubstring() {
	tests := []struct {
		s   string
		exp int
	}{
		{s: "abba", exp: 2},
	}

	for _, tt := range tests {
		got := longestSubstring(tt.s)
		fmt.Printf("exp: %d, got: %d\n", tt.exp, got)
	}
}
