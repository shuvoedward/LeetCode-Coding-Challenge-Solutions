package leetcode

func kthGrammar(n int, k int) int {
	if k == 1 {
		return 0
	} else if k%2 != 0 {
		return kthGrammar(n-1, (k+1)/2)
	} else {
		if kthGrammar(n-1, k/2) == 0 {
			return 1
		}
		return 0
	}
	/*
		779 note

		if k == 4, then the parent kth position is 2,
		\then from 2 its parent kth position is 1,
		starting at index 1, left child would always be odd,
		even child always would be on the right. so,
		if even flip it, if odd dont flip it,
	*/
}
