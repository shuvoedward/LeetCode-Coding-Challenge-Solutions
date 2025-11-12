package leetcode

// if curSum + nums[i] < num[i] then start the subarray from num[i]
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		curSum = max(nums[i], curSum+nums[i])

		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func MaxSubArrayAlt(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		maxSum = max(maxSum, curSum)
	}

	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
