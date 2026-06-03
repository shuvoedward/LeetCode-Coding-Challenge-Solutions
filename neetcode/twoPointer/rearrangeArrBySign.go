package twoPointer

import "fmt"

func rearrangeArray(nums []int) []int {
	// this does not maintain order
	for i, num := range nums {
		if i%2 == 0 {
			if num < 0 {
				j := i + 1
				for j < len(nums) && nums[j] < 0 {
					j++
				}
				nums[i], nums[j] = nums[j], nums[i]
			}
		} else {
			if num >= 0 {
				j := i + 1
				for j < len(nums) && nums[j] > 0 {
					j++
				}
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func rearrangeArray2(nums []int) []int {
	i, j := 0, 1
	res := make([]int, len(nums))

	for k := range len(nums) {
		if nums[k] > 0 {
			res[i] = nums[k]
			i += 2
		} else {
			res[j] = nums[k]
			j += 2
		}
	}

	return res
}

func TestRearrangeArr() {
	tests := []struct {
		nums []int
		exp  []int
	}{
		{nums: []int{3, 1, -2, -5, 2, -4}, exp: []int{3, -2, 1, -5, 2, -4}},
		{nums: []int{-1, 1}, exp: []int{1, -1}},
		{nums: []int{-1, -1, -1, 5, 5, 5}, exp: []int{5, -1, 5, -1, 5, -1}},
	}

	for _, tt := range tests {
		got := rearrangeArray(tt.nums)
		fmt.Printf("exp: %v, got: %v\n", tt.exp, got)
	}
}
