package leetcode

func minOperations(nums []int, x int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	target := total - x

	if target < 0 {
		return -1
	}

	if target == 0 {
		return len(nums)
	}

	sum := 0
	l := 0
	maxlen := -1

	for r, num := range nums {
		sum += num
		for sum > target && l <= r {
			sum -= nums[l]
			l++
		}

		if sum == target {
			maxlen = max(r-l+1, maxlen)
		}
	}

	if maxlen == -1 {
		return -1
	}

	return len(nums) - maxlen
}
