package leetcode

import "fmt"

func RemoveDuplicates(nums []int) int {
	// k := 0
	// count := map[int]bool{}

	// i, j := 0, 0

	// for j < len(nums) {
	// 	if !count[nums[j]] {
	// 		count[nums[j]] = true
	// 		k++
	// 		nums[i], nums[j] = nums[j], nums[i]
	// 		i++
	// 	}
	// 	j++
	// }
	// fmt.Println(nums)
	// return k
	var p int
	for i := 1; i < len(nums); i++ {
		if nums[p] != nums[i] {
			p++
		}
		nums[p] = nums[i]
	}
	fmt.Println(nums)
	return p + 1
}
