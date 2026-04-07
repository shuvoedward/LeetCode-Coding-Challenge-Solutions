package twoPointer

import (
	"unicode"
)

// make all character lowercase

func ValidPalindrome(s string) bool {
	// "Was it a car or a cat I saw?" true
	// A man, a plan, a canal: Panama => true

	// a man a plan a canal panama

	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphanumeric(rune(s[l])) {
			l++
		}
		for l < r && !isAlphanumeric(rune(s[r])) {
			r--
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphanumeric(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}
