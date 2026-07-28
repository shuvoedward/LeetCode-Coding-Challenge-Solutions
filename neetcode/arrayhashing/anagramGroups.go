package arrayhashing

func groupAnagrams(strs []string) [][]string {
	/*

	    [0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 1 0 0]: ["eat", "tea", "ate"],
	    [0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 1 0 0 0 0 0 1 0 0]: ["tan", "nat"],
	    [0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]: ["bat"]

		result := [][]string{
		it is a slice that its elements are slices of strings
	    {"eat", "tea", "ate"}, // First element, a slice of strings
	    {"tan", "nat"},        // Second element, another slice of strings
	    {"bat"},               // Third element, another slice of strings
	}


	*/
	res := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int

		for _, c := range s {
			count[c-'a']++
		}
		res[count] = append(res[count], s)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
	/*
	   const x = 3.14 type inference
	   var x int
	   var x int = 3.14
	   x:=3.14 // type inference

	   for slices
	   arr := []string {"strings", "strings"}
	   var arr []string = []string{"hello", "world"}
	   var arr = []string{"hello", "world"}

	   var arr []string
	   arr = []string {"hello", "world"}
	*/
}
