package leetcode

func FirstMissingPositive(nums []int) int {
	n := len(nums)

	// Step 1: Clean up - replace all numbers <= 0 or > n with n+1
	// These numbers are irrelevant since answer is in range [1, n+1]
	for i, num := range nums {
		if num <= 0 || num > n {
			nums[i] = n + 1
		}
	}

	// Step 2: Mark presence using negative signs
	// For each number x in range [1, n], mark nums[x-1] as negative
	for _, val := range nums {
		num := abs(val) // Use absolute value since it might already be negative

		// Only process if in valid range
		if num <= n {
			index := num - 1
			// Mark as negative (if not already)
			if nums[index] > 0 {
				nums[index] = -nums[index]
			}
		}
	}

	// Step 3: Find the first positive number
	// The first positive value at index i means (i+1) is missing
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}

	// If all numbers [1, n] are present, answer is n+1
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*

The key insight: The answer must be in the range [1, n+1] where n is the length of the array.

Why? Because if you have an array of length n, the smallest missing positive can't be larger than n+1. Think about it:

    Best case: array is [1, 2, 3, ..., n] → answer is n+1
    Any other case: answer is ≤ n

So we only care about numbers in the range [1, n]. Any number outside this range is irrelevant!

The strategy (using the array itself as a hash table):

    Ignore irrelevant numbers: Any number ≤ 0 or > n doesn't matter
    Use array indices as markers: Treat the array like a hash set where index i represents whether number i+1 exists
    Mark presence by making values negative: If we see number x, mark nums[x-1] as negative
    Find the first positive index: The first index that's still positive means that number is missing

The approach:

    Clean up the array: Replace all numbers ≤ 0 or > n with something harmless (like n+1)
    Mark presence: For each number x in range [1, n], mark nums[x-1] as negative (use absolute value since it might already be negative)
    Find missing: Scan through - the first positive value at index i means i+1 is missing

Example walkthrough with [3, 4, -1, 1]:

Array length n = 4, so we only care about [1, 2, 3, 4]

Step 1 - Clean: [3, 4, 5, 1] (replaced -1 with 5 since it's out of range)

Step 2 - Mark presence:

    See 3 → mark index 2 negative: [3, 4, -5, 1]
    See 4 → mark index 3 negative: [3, 4, -5, -1]
    See 5 → ignore (out of range)
    See 1 → mark index 0 negative: [-3, 4, -5, -1]

Step 3 - Find missing:

    Index 0: negative ✓ (1 exists)
    Index 1: positive ✗ (2 is missing!)

Return 2

Why this works: We're using the array's indices as a "presence checker" for numbers 1 through n. The negative sign is our mark

er without needing extra space!


*/
