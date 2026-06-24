package leetcode

import "fmt"

/*
Your task is to find the minimum number of subsequences from the source string that,
when concatenated together, form the target string.
If it's impossible to form the target string using any number of subsequences from source, return -1.

For example:

    If source = "abc" and target = "abcbc",
	you would need 2 subsequences: first taking "abc" and then taking "bc" from another
	iteration through source, giving you "abc" + "bc" = "abcbc".
    If source = "abc" and target = "acdbc",
	it would be impossible since 'd' doesn't exist in source, so you'd return -1.

*/

func findShortest(source, target string) int {
	subsequences := 0
	t := 0

	for t < len(target) {
		match := false
		s := 0
		for s < len(source) && t < len(target) {
			if source[s] == target[t] {
				t++
				match = true
			}
			s++
		}

		if !match {
			return -1
		}

		subsequences++
	}

	return subsequences
}

func TestFindShortest() {
	tests := []struct {
		source string
		target string
		exp    int
	}{
		{"abc", "abcbc", 2},
		{"abc", "abc", 1},
		{"abc", "abdbc", -1},
	}

	for _, tt := range tests {
		got := findShortest(tt.source, tt.target)
		fmt.Printf("exp: %d, got: %d\n", tt.exp, got)
	}
}

func shortestWay(source string, target string) int {
	// Create a map of character to indices for O(1) lookups
	charToIndices := make(map[rune][]int)
	for i, char := range source {
		charToIndices[char] = append(charToIndices[char], i)
	}

	// Check if all target characters exist
	for _, char := range target {
		if _, exists := charToIndices[char]; !exists {
			return -1
		}
	}

	subsequenceCount := 1
	currentPos := -1

	for _, char := range target {
		indices := charToIndices[char]

		// Binary search for next occurrence after currentPos
		left, right := 0, len(indices)-1
		nextPos := -1

		for left <= right {
			mid := (left + right) / 2
			if indices[mid] > currentPos {
				nextPos = indices[mid]
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		if nextPos == -1 {
			// Need new subsequence
			subsequenceCount++
			currentPos = indices[0]
		} else {
			currentPos = nextPos
		}
	}

	return subsequenceCount
}
