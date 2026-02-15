package hackerrank

func icecreamParlor(m int32, arr []int32) []int32 {
	// Write your code here
	seen := map[int32]int{}
	r := make([]int32, 2)

	for i, price := range arr {
		if price > m {
			continue
		}

		res := abs(m - price)

		if index, exists := seen[res]; exists {
			first, second := i+1, index
			if index < i {
				first = index
				second = i
			}
			r[0] = int32(first)
			r[1] = int32(second)

			return r
		}

		seen[price] = i + 1

	}

	return r
}

func abs(i int32) int32 {
	if i < 0 {
		return i * -1
	}
	return i
}
