package leetcode

func minAddToMakeValid(s string) int {
	open, close := 0, 0
	result := 0
	for _, p := range s {
		if p == '(' {
			open++
		} else {
			close = 1
		}

		open = open - close
		close = 0
		if open == -1 {
			result++
			open, close = 0, 0
		}
	}

	if open > 0 {
		result += open
	}

	return result
}

func minAddToMakeValidOptimized(s string) int {
	balance := 0
	addNeeded := 0
	for _, char := range s {
		if char == '(' {
			balance++
		} else {
			if balance > 0 {
				balance--
			} else {
				addNeeded++
			}
		}
	}
	return addNeeded + balance
}
