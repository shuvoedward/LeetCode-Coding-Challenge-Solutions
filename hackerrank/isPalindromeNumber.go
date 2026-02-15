package hackerrank

// ibm prep

func isPalindromeNumber(p int) bool {
	// negative numbers can not be palindrome
	if p < 0 {
		return false
	}
	// numbers ending with 0 are not palindrome, it has to start with zero as well
	// 20, 110
	if p%10 == 0 {
		return false
	}

	// even = 1221, 11
	// odd = 12321, 121
	// do this without turning it to strings
	reversedHalf := 0
	// for even, both will have reverse == p, 1221 rev = 12, p = 12
	// for odd reverseHalf > p, 121, reverses = 12, p = 1
	for reversedHalf < p {
		reversedHalf = reversedHalf*10 + (p % 10)
		p = p / 10
	}

	return reversedHalf == p || p == (reversedHalf/10)

}
