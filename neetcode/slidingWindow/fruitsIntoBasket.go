package slidingwindow

func totalFruits(fruits []int) int {
	count := make(map[int]int)
	l := 0

	res, total := 0, 0

	for r := range len(fruits) {
		count[fruits[r]]++
		total++

		for len(count) > 2 {
			total--
			count[fruits[l]]--
			if count[fruits[l]] == 0 {
				delete(count, fruits[l])
			}
			l++
		}

		if total > res {
			res = total
		}
	}

	return res
}
