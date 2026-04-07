package twoPointer

// always one solution
// no extra copy - no map, set
// no duplicate in the slice

func TwoSum(nums []int, target int) []int {
	// 1 + 2 = 5 > 3

	l, r := 0, len(nums)-1

	for l < r {
		sum := nums[l] + nums[r]

		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}

	}

	return []int{}
}

func HashmapTwosum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if val, exists := mp[tmp]; exists {
			return []int{val, i + 1}
		}
		mp[nums[i]] = i + 1
	}
	return []int{}
}
