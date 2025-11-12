package leetcode

func SearchRange(nums []int, target int) []int {
	left := findLeft(nums, target)
	right := findRight(nums, target)
	return []int{left, right}
}

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid    // Save this position
			right = mid - 1 // Keep searching LEFT
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid   // Save this position
			left = mid + 1 // Keep searching RIGHT
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

/*
## How It Works

### Finding Leftmost Position
When you find `target`:
- Save `mid` as a potential answer
- Search the **left half** (`right = mid - 1`)
- Keep updating `result` until you can't go left anymore

### Finding Rightmost Position
When you find `target`:
- Save `mid` as a potential answer
- Search the **right half** (`left = mid + 1`)
- Keep updating `result` until you can't go right anymore

## Example Walkthrough

`nums = [5,7,7,8,8,8,10]`, `target = 8`

### findLeft(nums, 8):
```
Iteration 1: left=0, right=6, mid=3
nums[3]=8 == target
result=3, search left → right=2

Iteration 2: left=0, right=2, mid=1
nums[1]=7 < 8
search right → left=2

Iteration 3: left=2, right=2, mid=2
nums[2]=7 < 8
search right → left=3

left > right, stop
Return result=3 (leftmost position)
```

### findRight(nums, 8):
```
Iteration 1: left=0, right=6, mid=3
nums[3]=8 == target
result=3, search right → left=4

Iteration 2: left=4, right=6, mid=5
nums[5]=8 == target
result=5, search right → left=6

Iteration 3: left=6, right=6, mid=6
nums[6]=10 > 8
search left → right=5

left > right, stop
Return result=5 (rightmost position)

*/
