package hackerrank

// https://www.hackerrank.com/challenges/two-characters/problem?isFullScreen=true

func TwoCharacters(s string) int32 {
	b := []byte(s)

	present := make([]bool, 26)
	for _, ch := range b {
		present[ch-'a'] = true
	}

	best := 0

	for i := 0; i < 26; i++ {
		if !present[i] {
			continue
		}
		for j := i + 1; j < 26; j++ {
			if !present[j] {
				continue
			}

			var last byte = 0
			length := 0
			valid := true

			// convert the characters back
			ci := byte('a' + i)
			cj := byte('a' + j)

			for _, ch := range b {
				if ch != ci && ch != cj {
					continue
				}

				if length > 0 && ch == last {
					valid = false
					break
				}

				last = ch
				length++
			}

			if valid && length > best {
				best = length
			}
		}
	}

	return int32(best)
}
