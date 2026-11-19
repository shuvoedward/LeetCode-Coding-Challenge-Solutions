package binarysearch

import "fmt"

func findTarget(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		}

		if nums[l] < nums[m] {
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}

	}

	return -1
}

func Test_findTarget() {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{4, 1, 2, 3}, target: 3, expected: 3},
	}

	for i, tt := range tests {
		output := findTarget(tt.nums, tt.target)
		fmt.Printf("%d. expected: %d, got: %d", i, tt.expected, output)
	}

}
