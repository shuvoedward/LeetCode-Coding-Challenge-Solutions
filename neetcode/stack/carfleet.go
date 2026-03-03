package stack

import "sort"

func CarFleet(target int, position []int, speed []int) int {
	// use array inside slice, the size is know and array is faster than slice
	pair := make([][2]int, len(position))

	// for slice, declaring size and indexing is faster than append
	for i := 0; i < len(position); i++ {
		pair[i] = [2]int{position[i], speed[i]}
	}

	// sort based on position, in descending order
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] > pair[j][0]
	})

	// tracks the fleet
	stack := []float64{}

	for _, p := range pair {
		time := float64(target-p[0]) / float64(p[1])
		stack = append(stack, time)
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}

/*
Time complexity: O(nlogn) - because of sorting
Space complexity: O(n) - pair and stack slice
*/
