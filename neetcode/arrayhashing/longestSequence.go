package arrayhashing

// Brute force
// Time Complexity n^2
// Space n
func SequenceBruteForce(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// This is golang way of using set similar to js.
	// value is not needed. And empty struct does not allocate any memory
	store := make(map[int]struct{})
	res := 0

	for _, num := range nums {
		store[num] = struct{}{} // initialize empty struct, allocates no memory
	}
	// 2, 20, 4, 10, 3, 4, 5
	// range works with iterable types like slices, arrays, maps, strings and channels
	for num := range store {
		// num only holds the key, to hold value, another variable needs to be provided in for loop
		// for num, value := range store
		// Start only if num is the beginning of a sequence
		if _, found := store[num-1]; !found {
			streak, curr := 1, num

			for _, ok := store[curr+1]; ok; _, ok = store[curr+1] {
				streak++
				curr++
			}

			if streak > res {
				res = streak
			}
		}
	}
	return res
}

func SequenceHashmap(nums []int) int {

	// n Go, when you access a map with a key that does not exist,
	// it returns the zero value for the value type. Since the map
	// value type is int, the default value is 0.

	//nums = [100, 4, 200, 1, 3, 2, 4]
	// mp = {100:1, 4:4, 200:1, 1:4, 3:4, 2:4}

	mp := make(map[int]int)
	res := 0

	for _, num := range nums {
		if mp[num] == 0 { // Checks for duplicates
			left := mp[num-1] // If left and right, or either exists then num is part of a sequence
			right := mp[num+1]
			// left and right, holds the position of the that number in the sequence.

			sum := left + right + 1 // 1 is for the current num

			mp[num] = sum
			// Updating the boundaries
			mp[num-left] = sum  // leftmost number, keeps track of the sequence size)
			mp[num+right] = sum // rightmost number, keeps track of the sequence size)

			if sum > res {
				res = sum
			}
		}
	}
	return res
}

func SequenceSet(nums []int) int {
	// go's way to create set, struct does not allocates any memory if no value initalized
	//
	numSet := make(map[int]struct{})

	for _, num := range nums {

		numSet[num] = struct{}{}
	}

	// Set look up is constant time. and element based look up not index based.

	//nums = [100, 4, 200, 1, 3, 2, 4]
	longest := 0
	for num := range numSet { // ignores the value, uses only the key
		// is num a start of a sequence?
		if _, found := numSet[num-1]; !found {
			length := 1

			for {
				if _, exists := numSet[num+length]; exists {
					length++
				} else {
					break
				}
			}

			if longest < length {
				longest = length
			}

		}
	}
	return longest
}

func longestConsecutive(nums []int) int {
	mp := map[int]bool{}
	res := 0

	for _, num := range nums {
		mp[num] = true
	}

	for _, num := range nums {
		curMax := 1
		i := 1
		if mp[num-1] {
			continue
		}

		for mp[num+i] {
			curMax++
			i++
		}

		if curMax > res {
			res = curMax
		}
	}

	return res
}
