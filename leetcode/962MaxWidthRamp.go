package leetcode

func maxWidthRamp(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] <= nums[r] {
			return r - l
		}
		if nums[l] > nums[r] {
			l++
		}
	}

	return 0
}
