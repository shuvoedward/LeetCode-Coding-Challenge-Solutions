package leetcode

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{}, len(wordDict))
	maxLen := 0
	for _, w := range wordDict {
		dict[w] = struct{}{}
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		// Only need to look back up to maxLen characters
		start := i - maxLen
		if start < 0 {
			start = 0
		}
		for j := start; j < i; j++ {
			if !dp[j] {
				continue
			}
			if _, ok := dict[s[j:i]]; ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func RunWordBreakTest() {
	tests := []struct {
		s        string
		wordDict []string
		exp      bool
	}{
		// {s: "leetcode", wordDict: []string{"leet", "code"}, exp: true},
		// {s: "applepenapple", wordDict: []string{"apple", "pen"}, exp: true},
		// {s: "bb", wordDict: []string{"a", "b", "bbb", "bbbb"}, exp: true},
		// {s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}, exp: false},
		// {s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}, exp: false},
		{s: "abcd", wordDict: []string{"a", "abc", "b", "cd"}, exp: false},
	}

	for _, tt := range tests {
		got := wordBreak(tt.s, tt.wordDict)
		fmt.Printf("exp: %t, got: %t\n", tt.exp, got)
	}
}
