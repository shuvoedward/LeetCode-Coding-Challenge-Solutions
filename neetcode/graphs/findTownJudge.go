package graphs

func findJudge(n int, trusts [][]int) int {
	in := make(map[int]int)
	out := make(map[int]int)

	for _, t := range trusts {
		out[t[0]]++
		in[t[1]]++
	}

	for person, incoming := range in {
		if incoming == n-1 && out[person] == 0 {
			return person
		}
	}

	return -1
}
