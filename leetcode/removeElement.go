package leetcode

import "fmt"

func RemoveElement(nums []int, val int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[l] == val {
			nums[l] = nums[r]
			r--
		} else {
			l++
		}
	}

	fmt.Println(nums)
	return l
}
