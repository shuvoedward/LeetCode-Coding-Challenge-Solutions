package hackerrank

import "fmt"

func WeightedUniformStrings(s string, queries []int32) []string {
	// find weight of the substrings
	// insert each weight in map
	// find queries inside the map

	b := []byte(s)
	weights := map[int]bool{}

	curWeight := 0
	prev := b[0]

	for _, ch := range b {
		w := int(ch-'a') + 1

		if ch == prev {
			curWeight += w
		} else {
			curWeight = w
			prev = ch
		}
		weights[curWeight] = true
	}

	res := make([]string, 0, len(queries))

	for i := 0; i < len(queries); i++ {
		if weights[int(queries[i])] {
			res = append(res, "Yes")
		} else {
			res = append(res, "No")
		}

	}

	return res
}

func RunWeightUniformStringsTest() {
	test := []struct {
		s       string
		queries []int32
		exp     []string
	}{
		{s: "aaabbbbcccddd", queries: []int32{5, 9, 7, 8, 12, 5}, exp: []string{"Yes", "No", "Yes", "Yes", "No"}},
		{s: "abccddde", queries: []int32{6, 1, 3, 12, 5, 9, 10}, exp: []string{"Yes", "yes", "Yes", "Yes", "No", "No"}},
	}

	for _, tt := range test {
		res := WeightedUniformStrings(tt.s, tt.queries)
		fmt.Printf("exp: %v, got: %v\n", tt.exp, res)
	}
}
