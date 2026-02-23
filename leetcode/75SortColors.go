package leetcode

func sortColors(nums []int) {
	// sort in place, dont return the array, don't use sort function
	// 0 = red
	// 1 = white
	// 2 = blue
	// sort in place in ascending order but in range of 0, 1, 2
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

/*

low = stays at the position where 0 can be inserted
mid = stays at the position where 1 can be inserted
high = stays at the position where 2 can be inserted

*/
