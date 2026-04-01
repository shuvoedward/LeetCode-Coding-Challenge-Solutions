package leetcode75

import (
	"fmt"
	"slices"
)

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

func closeStringsOptimized(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := [26]int{}
	freq2 := [26]int{}

	for i := range len(word1) {
		freq1[word1[i]-'a']++
		freq2[word2[i]-'a']++
	}

	for i := range 26 {
		if (freq1[i] == 0) != (freq2[i] == 0) {
			return false
		}
	}

	slices.Sort(freq1[:])
	slices.Sort(freq2[:])

	for i := range 26 {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}

func test() {
	tests := []struct {
		word1    string
		word2    string
		expected bool
	}{
		{"aba", "baa", true},
		{"abcab", "dabca", true},
		{"a", "b", false},
		{"cabbac", "abbaca", true},
	}

	for _, tt := range tests {
		result := CloseStrings(tt.word1, tt.word2)
		fmt.Printf("closeStrings(\"%s\", \"%s\") = %v (expected %v)\n",
			tt.word1, tt.word2, result, tt.expected)
	}
}
