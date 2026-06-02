package slidingwindow

func decrypt(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)

	if k == 0 {
		return result
	}

	// Compute initial window sum
	windowSum := 0
	l, r := 0, 0

	if k > 0 {
		l, r = 1, k
	} else {
		l, r = n+k, n-1
	}

	for i := l; i <= r; i++ {
		windowSum += code[i%n]
	}

	for i := 0; i < n; i++ {
		result[i] = windowSum
		// Slide window: add new element entering, remove element leaving
		windowSum += code[(r+1)%n]
		windowSum -= code[l%n]
		l++
		r++
	}

	return result
}

func decrypt2(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)

	if k == 0 {
		return result
	}

	for i := 0; i < n; i++ {
		sum := 0
		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%n]
			}
		} else {
			for j := 1; j <= -k; j++ {
				sum += code[(i-j+n)%n]
			}
		}
		result[i] = sum
	}

	return result
}
