package arrayhashing

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)

	// If ch is not already a key in the map,
	// Go initializes the value for ch to its zero value
	//(0) automatically. Then it increments it to 1.
	for i, ch := range s {
		countS[ch]++
		// t[i] gets the character at index i of string t.
		// rune(t[i]) ensures that the character is treated as a rune for the countT map (although it’s not strictly necessary here since t[i] is already convertible to rune).
		//countT[rune(t[i])]++ increments the count of the corresponding character in the countT map.
		countT[rune(t[i])]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}

	return true

}

func isAnagramSingleArr(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}

	for i := range len(s) {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true
}
