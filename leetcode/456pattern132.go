package leetcode

func find132pattern(nums []int) bool {
	for k := 2; k < len(nums); k++ {
		if nums[k-2] < nums[k] && nums[k] < nums[k-1] {
			return true
		}
	}

	return false
}
