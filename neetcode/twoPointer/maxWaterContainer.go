package twoPointer

func maxArea(height []int) int {
	l, r := 0, len(height)-1

	result := 0

	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > result {
			result = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
