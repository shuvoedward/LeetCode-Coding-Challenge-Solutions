package leetcode

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			threeSum := a + nums[left] + nums[right]
			if threeSum < 0 {
				left++
			} else if threeSum > 0 {
				right--
			} else {
				res = append(res, []int{a, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return res
}

func RunTestThreeSum() {
	tests := []struct {
		nums []int
		exp  [][]int
	}{
		{nums: []int{-1, 0, 1, 2, -1, -4}, exp: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{nums: []int{1, 2, 0, 1, 0, 0, 0, 0}, exp: [][]int{{0, 0, 0}}},
		{nums: []int{0, 0, 0}, exp: [][]int{{0, 0, 0}}},
	}

	fmt.Println("Three sum tests")
	for _, tt := range tests {
		got := threeSum(tt.nums)
		fmt.Printf("exp: %v, got: %v\n", tt.exp, got)
	}
}
