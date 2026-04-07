package twoPointer

// Hard

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	res := 0

	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			res += leftMax - height[l]
		} else {
			r--
			rightMax = max(rightMax, height[r])
			res += rightMax - height[r]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func trapPreSufArr(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	res := 0
	for i := range n {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}

	return res
}
