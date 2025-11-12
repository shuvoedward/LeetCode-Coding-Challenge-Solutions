package leetcode

func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Determine which half is sorted
		if nums[left] <= nums[mid] {
			// Left half is sorted
			if nums[left] <= target && target < nums[mid] {
				// Target is in the sorted left half
				right = mid - 1
			} else {
				// Target is in the right half
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if nums[mid] < target && target <= nums[right] {
				// Target is in the sorted right half
				left = mid + 1
			} else {
				// Target is in the left half
				left = mid - 1
			}
		}
	}

	return -1
}
