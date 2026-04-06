package binarysearch

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x

	res := 0

	for l <= r {
		mid := l + (r-l)/2

		if int64(mid)*int64(mid) > int64(x) {
			r = mid - 1
		} else if int64(mid)*int64(mid) < int64(x) {
			l = mid + 1
			res = mid
		} else {
			return mid
		}
	}

	return res
}

func TestMySqrt() {
	tests := []struct {
		x   int
		exp int
	}{
		{x: 25, exp: 5},
		{x: 125, exp: 5},
	}

	for _, tt := range tests {
		fmt.Printf("exp: %d, got: %d\n", tt.exp, mySqrt(tt.x))
	}
}
