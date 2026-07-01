package graphs

func isAlienSorted(words []string, order string) bool {
	orderIndex := make([]int, 26)
	for i, ch := range order {
		orderIndex[ch-'a'] = i
	}

	for i := range len(words) - 1 {
		w1, w2 := words[i], words[i+1]

		for j := range len(w1) {
			if j == len(w2) {
				return false
			}

			if w1[j] != w2[j] {
				if orderIndex[w1[j]-'a'] > orderIndex[w2[j]-'a'] {
					return false
				}
				break
			}
		}
	}

	return true
}
