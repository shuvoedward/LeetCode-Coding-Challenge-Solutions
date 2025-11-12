package binarysearch

func FindMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2
		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
