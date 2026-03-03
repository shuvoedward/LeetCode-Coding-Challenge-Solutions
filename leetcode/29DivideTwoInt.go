package leetcode

import (
	"fmt"
	"math"
)

func divideTwoInt(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	neg := (dividend < 0) != (divisor < 0) // return true if both are not equal, otherwise false
	// true = true != false
	// false = true != true

	a := int64(dividend)
	b := int64(divisor)
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	var res int64
	for a >= b {
		temp := b
		multiple := int64(1)

		for (temp << 1) <= a {
			temp = temp << 1
			multiple = multiple << 1
		}

		a -= temp
		res += multiple
	}

	if neg {
		res = -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}

	return int(res)
}

func RunDivideTwoIntTest() {
	tests := []struct {
		dividend int
		divisor  int
		exp      int
	}{
		{dividend: 7, divisor: -3, exp: -2},
	}

	for _, tt := range tests {
		got := divideTwoInt(tt.dividend, tt.divisor)
		fmt.Printf("exp: %d, got: %d\n", tt.exp, got)
	}
}
