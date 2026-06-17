package leetcode

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num))

	for i := 0; i < len(num); i++ {
		digit := num[i]
		// pop while there's a bigger digit sitting on top of a smaller one
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > digit {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}

	// leftover removals (e.g. "111", k=2) — trim from the end
	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	// strip leading zeros
	i := 0
	for i < len(stack)-1 && stack[i] == '0' {
		i++
	}
	stack = stack[i:]

	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
