package leetcode

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func myAtoI(s string) int {
	result := 0
	isNeg := false

	leading := false
	var curNumStr strings.Builder
	for _, ch := range s {
		if ch == ' ' {
			continue
		} else if len(curNumStr.String()) == 0 && (ch == '-' || ch == '+') {
			if leading {
				break
			}
			leading = true
			if ch == '-' {
				isNeg = true
			} else {
				continue
			}
		} else if unicode.IsDigit(ch) {
			if ch == '0' && len(curNumStr.String()) == 0 {
				leading = true
				continue
			}
			curNumStr.WriteRune(ch)
		} else {
			break
		}
	}

	numMap := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	const INT_MAX = math.MaxInt32
	const INT_MIN = math.MinInt32

	extractedStr := curNumStr.String()
	strLen := len(extractedStr)
	for _, ch := range extractedStr {
		num := numMap[ch]
		if !isNeg {
			if result > (INT_MAX - num) {
				return INT_MAX
			}
		} else {
			if result > (-(INT_MIN + num))/10 {
				return INT_MIN
			}
		}
		result = result*10 + num
		strLen--

	}
	if isNeg {
		result *= -1
	}

	return result
}

func RunMyAtoITest() {
	tests := []struct {
		s   string
		exp int
	}{
		// {s: "42", exp: 42},
		// {s: "     -042", exp: -42},
		// {s: "1337c0d3", exp: 1337},
		// {s: "0-1", exp: 0},
		{s: "+1", exp: 1},
		{s: "+-12", exp: 0},
	}

	for _, tt := range tests {
		got := myAtoI(tt.s)
		fmt.Printf("exp: %d, got: %d\n", tt.exp, got)
	}
}
