package leetcode75

import "fmt"

func CloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := map[rune]int{}
	freq2 := map[rune]int{}

	for _, ch := range word1 {
		freq1[ch]++
	}

	for _, ch := range word2 {
		freq2[ch]++
	}

	for ch := range freq1 {
		if freq2[ch] == 0 {
			return false
		}
	}

	for ch := range freq2 {
		if freq1[ch] == 0 {
			return false
		}
	}

	freqCount1 := map[int]int{}
	freqCount2 := map[int]int{}

	for _, count := range freq1 {
		freqCount1[count]++
	}
	for _, count := range freq2 {
		freqCount2[count]++
	}

	if len(freqCount1) != len(freqCount2) {
		return false
	}

	for key, val := range freqCount1 {
		if freqCount2[key] != val {
			return false
		}
	}

	return true
}

func test() {
	testCases := []struct {
		word1    string
		word2    string
		expected bool
	}{
		{"aba", "baa", true},
		{"abcab", "dabca", true},
		{"a", "b", false},
		{"cabbac", "abbaca", true},
	}

	for _, tc := range testCases {
		result := CloseStrings(tc.word1, tc.word2)
		fmt.Printf("closeStrings(\"%s\", \"%s\") = %v (expected %v)\n",
			tc.word1, tc.word2, result, tc.expected)
	}
}
