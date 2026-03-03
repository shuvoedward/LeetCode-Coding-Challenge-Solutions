package leetcode

import "fmt"

func findDuplicate(nums []int) int {
	slow := 0
	fast := 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			return slow
		}
	}

}

func RunFindDuplicateTest() {
	tests := []struct {
		nums []int
		exp  int
	}{
		{nums: []int{1, 3, 4, 2, 2}, exp: 2},
		{nums: []int{3, 1, 3, 4, 2}, exp: 3},
	}

	for _, tt := range tests {
		got := findDuplicate(tt.nums)
		fmt.Printf("exp: %d, got: %d\n", tt.exp, got)
	}
}
