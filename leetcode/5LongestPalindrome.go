package leetcode

func longestPalindrome(s string) string {
	// babad
	start, end := 0, 0

	for i := range s {
		len1 := expand(s, i, i)
		len2 := expand(s, i, i+1)

		maxLen := max(len1, len2)

		if maxLen > end-start {
			// i is middle. so, using maxlength and middle index, how to
			// know the start and end of the substring?
			// because i == middle,
			// divide maxlen / 2,
			// for after middle, i + maxLen/2
			// for before i - maxlen/2
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left + 1
}
